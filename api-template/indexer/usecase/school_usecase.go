package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/consts"
	"github.com/MStation-io/api/contracts/school"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/repo"
	"github.com/MStation-io/api/repo/pg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	UpgradeLevelEvent = "UpgradeLevel(address,uint256,bool,uint16)"
)

var schoolSupportedLogs = []string{
	UpgradeLevelEvent,
}

var schoolKeccakMap map[string]common.Hash = map[string]common.Hash{}

type upgradeCharacterLog struct {
	NftAddress   string
	OwnerAddress string
	TokenId      string
	Result       bool
	NewLevel     int
	TxHash       string
	BlockNumber  int
}

type schoolContract struct {
	address  common.Address
	instance *school.School
}

type SchoolUsecase struct {
	cfg      config.Config
	store    repo.DBRepo
	repo     *repo.Repo
	client   *ethclient.Client
	contract schoolContract
}

type upgradeCharacterlogData struct {
	Result    string `json:"Result"`
	FromLevel int    `json:"FromLevel"`
	NewLevel  int    `json:"NewLevel"`
}

func NewSchoolUsecase(cfg config.Config, store repo.DBRepo, client *ethclient.Client) IUsecase {
	schoolContractAddress := common.HexToAddress(cfg.SchoolContract)
	school, err := school.NewSchool(schoolContractAddress, client)
	if err != nil {
		zap.L().Fatal("cannot init school contract", zap.Error(err))
	}
	repo := pg.NewRepo()

	return SchoolUsecase{
		cfg:    cfg,
		store:  store,
		repo:   repo,
		client: client,

		contract: schoolContract{
			address:  schoolContractAddress,
			instance: school,
		},
	}

}

func (uc SchoolUsecase) Process(log types.Log) error {
	// switch case
	switch log.Topics[0].String() {
	case schoolKeccakMap[UpgradeLevelEvent].String():
		param, err := uc.upgradeCharacterFromLog(log)
		if err != nil {
			zap.L().Error("[usecase.Process] uc.upgradeCharacterFromLog()")
			return err
		}
		return uc.handleUpgradeCharacter(param)
	}
	return nil
}

func (uc SchoolUsecase) Topics() []common.Hash {
	hashes := make([]common.Hash, 0)
	for _, l := range schoolSupportedLogs {
		hash := common.BytesToHash(crypto.Keccak256([]byte(l)))
		hashes = append(hashes, hash)
		schoolKeccakMap[l] = hash
	}
	return hashes
}

func (uc SchoolUsecase) ShouldProcessLog(log types.Log) bool {
	if log.Address.Hex() != uc.contract.address.Hex() {
		return false
	}
	for _, hash := range schoolKeccakMap {
		if log.Topics[0] == hash {
			return true
		}
	}
	return false
}

func (uc SchoolUsecase) Name() string {
	return "SchoolUsecase"
}

func (uc SchoolUsecase) ContractAddress() common.Address {
	return uc.contract.address
}

func (uc SchoolUsecase) handleUpgradeCharacter(param *upgradeCharacterLog) error {

	ctx, cancel := context.WithTimeout(context.Background(), consts.DefaultLockTime*time.Second)
	defer cancel()

	tx, done := uc.store.NewTransactionWithContext(ctx)

	result := "Failed"
	if param.Result {
		result = "Success"
	}

	dataLogRaw := upgradeCharacterlogData{
		Result:    result,
		FromLevel: 1,
		NewLevel:  param.NewLevel,
	}

	data, err := json.Marshal(dataLogRaw)
	if err != nil {
		zap.L().Error("[uc.handleUpgradeCharacter] json.Marshal()",
			zap.Error(err),
			zap.Any("block_number", param.BlockNumber))
		return done(err)
	}

	if err != nil {
		zap.L().Error("[uc.handleNewBreed] uc.repo.CharacterLog.Create()",
			zap.Error(err),
			zap.Any("block_number", param.BlockNumber))
		return done(err)
	}

	return done(nil)
}

func (uc SchoolUsecase) upgradeCharacterFromLog(log types.Log) (*upgradeCharacterLog, error) {

	eventLog, err := uc.contract.instance.ParseUpgradeLevel(log)
	if err != nil {
		zap.L().Error("[uc.upgradeCharacterFromLog] uc.contract.instance.ParseUpgradeLevel(): ", zap.Error(err))
		return nil, err
	}

	tx, _, err := uc.client.TransactionByHash(context.Background(), common.HexToHash(eventLog.Raw.TxHash.Hex()))
	if err != nil {
		zap.L().Error("e.client.TransactionByHash(): ", zap.Error(err))
		return nil, errors.ErrInvalidTxHash
	}
	chainId, err := uc.client.NetworkID(context.Background())
	if err != nil {
		zap.L().Fatal("cannot get chainId", zap.Error(err))
	}
	signer := types.LatestSignerForChainID(chainId)

	from, err := types.Sender(signer, tx)
	if err != nil {
		zap.L().Error("types.Sender(): ", zap.Error(err))
		return nil, errors.ErrInternalServerError
	}

	return &upgradeCharacterLog{
		NftAddress:   eventLog.NftAddress.Hex(),
		OwnerAddress: from.Hex(),
		TokenId:      eventLog.TokenId.String(),
		Result:       eventLog.Result,
		NewLevel:     int(eventLog.NewLevel),
		TxHash:       eventLog.Raw.TxHash.Hex(),
		BlockNumber:  int(eventLog.Raw.BlockNumber),
	}, nil
}
