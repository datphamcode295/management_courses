package main

import (
	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/indexer"
	"github.com/MStation-io/api/log"
	"github.com/MStation-io/api/repo/pg"
	"go.uber.org/zap"
)

func main() {
	//init config
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)

	// init logger
	undo := log.New()
	defer zap.L().Sync()
	defer undo()

	s, close := pg.NewPostgresStore(&cfg)
	defer close()

	idxSvc, err := indexer.NewIndexService(cfg, s)
	if err != nil {
		zap.L().Panic("cannot init indexer", zap.Error(err))
	}

	idxSvc.Index()
}
