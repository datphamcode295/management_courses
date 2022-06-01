package indexer

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/indexer/usecase"
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
	"github.com/MStation-io/api/repo/pg"
	"github.com/MStation-io/api/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type IndexService struct {
	cfg      config.Config
	store    repo.DBRepo
	client   *ethclient.Client
	tipBH    int64
	repo     *repo.Repo
	usecases []usecase.IUsecase
}

func NewIndexService(cfg config.Config, store repo.DBRepo) (*IndexService, error) {
	// init a client instance
	client, err := ethclient.Dial(cfg.BscRPCEndPoint)
	if err != nil {
		zap.L().Error("cannot init client", zap.Error(err))
		return nil, err
	}
	repo := pg.NewRepo()

	// init usecase here
	usecases := make([]usecase.IUsecase, 0)
	usecases = append(usecases, usecase.NewSchoolUsecase(cfg, store, client))

	processedBlock := getLastestProcessedBlock(store, repo, cfg.IndexFromBlock)

	// start fetch from latest + 1
	bh := processedBlock.BlockNumber + 1

	return &IndexService{
		cfg:      cfg,
		store:    store,
		client:   client,
		usecases: usecases,
		repo:     repo,
		tipBH:    bh,
	}, nil
}

func (svc *IndexService) Index() {
	filterAddress := []common.Address{}
	for _, v := range svc.usecases {
		filterAddress = append(filterAddress, v.ContractAddress())
	}
	filterTopics := svc.getTopicsFromUsesCases()
	for {

		// first, we check block generation by calling blockheader
		zap.L().Info("fetching block header", zap.Any("seq", svc.tipBH))
		lbh, err := svc.client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			zap.L().Warn("cannot fetch latest block header", zap.Error(err), zap.Any("seq", svc.tipBH))
			time.Sleep(10 * time.Second)
			continue
		}

		bh, err := svc.client.HeaderByNumber(context.Background(), big.NewInt(svc.tipBH))
		if err != nil {
			zap.L().Warn("cannot fetch tip block header", zap.Error(err), zap.Any("seq", svc.tipBH))
			time.Sleep(10 * time.Second)
			continue
		}

		/*
			Handle relative confident finality: practice worst case scenario of
			expecting an evil set of 21 validators to do bad things.
			https://legendfantasywar.atlassian.net/wiki/spaces/LOFW/pages/6193153/Backend+scalability+outstanding+issues

			BNB average block time >= 3s
			Best case finality = ((2/3) * 21 + 1) * 3 = 45s
			Best case buffered blocks to account for = 45 / 3 = 15 blocks

			Worst case finality = ((2/3) * 21 + 1) * 4 = 60s (we will take 4s to account for drift)
			Worst case buffered blocks to account for = 60 / 3 = 20 blocks
		*/
		var behindTime int64 = 3 * 20
		timeOffset := time.Now().Unix() - int64(bh.Time)
		blockOffset := lbh.Number.Int64() - bh.Number.Int64()

		timeTooSoon := timeOffset < behindTime
		blockTooSoon := blockOffset < 20

		if timeTooSoon || blockTooSoon {
			zap.L().Debug("wait for finality", zap.Any("seq", svc.tipBH), zap.Any("block_offset", blockOffset))
			// Prevent a zero or negative sleep offset and retry
			sleepOffset := behindTime - timeOffset
			boundSleepOffset := util.Max(sleepOffset, 3)
			time.Sleep(time.Duration(boundSleepOffset) * time.Second)
			continue
		}

		// get all logs based on usecases's needed
		logs, err := svc.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: bh.Number,
			ToBlock:   bh.Number,
			Topics: [][]common.Hash{
				filterTopics,
			},
			Addresses: filterAddress,
		})
		if err != nil {
			zap.L().Fatal("cannot get logs", zap.Error(err))
			continue
		}

		zap.L().Info("processing logs ...", zap.Any("len", len(logs)), zap.Any("block_number", bh.Number))
	OUTER:
		for _, l := range logs {
			for _, usecase := range svc.usecases {
				if usecase.ShouldProcessLog(l) {
					zap.L().Info(fmt.Sprintf("process by %s", usecase.Name()), zap.Any("txHash", l.TxHash), zap.Any("topics[0]", l.Topics[0]))
					err := usecase.Process(l)
					if err != nil {
						zap.L().Fatal("panic when process block", zap.Any("debug data", l))
					}
					continue OUTER
				}
			}
		}

		_, err = svc.repo.ProcessedBlock.UpdateBlockNumberByBlockNumber(svc.store, svc.tipBH-1, bh.Number.Int64())
		if err != nil {
			zap.L().Error("cannot update processed_block", zap.Error(err))
			return
		}

		svc.tipBH++
	}
}

func (svc *IndexService) getTopicsFromUsesCases() []common.Hash {
	topics := make([]common.Hash, 0)
	for i := range svc.usecases {
		topics = append(topics, svc.usecases[i].Topics()...)
	}
	return topics
}

func getLastestProcessedBlock(store repo.DBRepo, repo *repo.Repo, blockNumber string) model.ProcessedBlock {
	processedBlock, err := repo.ProcessedBlock.GetCurrentProcessedBlock(store)
	if err != nil {
		zap.L().Error("cannot get processed_block", zap.Error(err))
		zap.L().Info("use default INDEX_FROM_BLOCK env", zap.Any("value", blockNumber))

		var ferr error
		processedBlock.BlockNumber, ferr = strconv.ParseInt(blockNumber, 10, 64)
		if ferr != nil {
			zap.L().Fatal("cannot get INDEX_FROM_BLOCK env", zap.Error(err))
		}
		processedBlock, err = repo.ProcessedBlock.Create(store, *processedBlock)
		if err != nil {
			zap.L().Fatal("cannot create processedBlock", zap.Error(err))
		}
	}
	return *processedBlock
}
