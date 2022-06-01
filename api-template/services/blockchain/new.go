package blockchain

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type blockchain struct {
	store       repo.DBRepo
	repo        *repo.Repo
	cfg         config.Config
	blockNumber int64
	client      *ethclient.Client
	chainId     *big.Int
}

func NewBlockChainService(cfg config.Config, store repo.DBRepo, repo *repo.Repo) BlockChainService {
	client, err := ethclient.Dial(cfg.BscRPCEndPoint)
	if err != nil {
		zap.L().Fatal("cannot init blockchain", zap.Error(err))
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		zap.L().Fatal("cannot get chainId", zap.Error(err))
	}

	blockNumber, _ := strconv.Atoi(cfg.IndexFromBlock)

	service := &blockchain{
		blockNumber: int64(blockNumber),
		store:       store,
		repo:        repo,
		cfg:         cfg,
		client:      client,
		chainId:     chainId,
	}
	go service.SyncCurrentHeadBlock()
	return service
}

func (b *blockchain) CurrentHeadBlock() int64 {
	return b.blockNumber
}

func (b *blockchain) SyncCurrentHeadBlock() {
	for {
		header, err := b.client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			zap.L().Warn("cannot fetch latest block header", zap.Error(err))
			continue
		}
		b.blockNumber = header.Number.Int64()
		time.Sleep(time.Second * 1)
	}
}
