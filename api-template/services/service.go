package services

import (
	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/repo"
)

type Services struct {
	// BlockChain blockchain.BlockChainService
	// Monitor monitor.MonitorService
}

func NewServices(cfg config.Config, store repo.DBRepo, repo *repo.Repo) *Services {
	return &Services{
		// BlockChain: blockchain.NewBlockChainService(cfg, store, repo),
		// Monitor: monitor.NewMonitorService(cfg, store, repo),
	}
}
