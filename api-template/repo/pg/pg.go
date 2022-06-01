package pg

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/repo"
	"github.com/MStation-io/api/util"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

// store is implimentation of repository
type store struct {
	database *gorm.DB
}

// DB database connection
func (s *store) DB() *gorm.DB {
	return s.database
}

// NewTransaction for database connection
// return an db instance and a done(error) function, if the error != nil -> rollback
func (s *store) NewTransaction() (newRepo repo.DBRepo, finallyFn repo.FinallyFunc) {
	newDB := s.database.Begin()
	finallyFn = func(err error) error {
		if err != nil {
			nErr := newDB.Rollback().Error
			if nErr != nil {
				return errors.NewStringError(nErr.Error(), http.StatusInternalServerError)
			}
			return errors.NewStringError(err.Error(), util.ParseErrorCode(err))
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return errors.NewStringError(cErr.Error(), http.StatusInternalServerError)
		}
		return nil
	}

	return &store{database: newDB}, finallyFn
}

func (s *store) NewTransactionWithContext(ctx context.Context) (newRepo repo.DBRepo, finallyFn repo.FinallyFunc) {
	newDB := s.database.WithContext(ctx).Begin()
	finallyFn = func(err error) error {
		if err != nil {
			nErr := newDB.Rollback().Error
			if nErr != nil {
				return errors.NewStringError(nErr.Error(), http.StatusInternalServerError)
			}
			return errors.NewStringError(err.Error(), util.ParseErrorCode(err))
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return errors.NewStringError(cErr.Error(), http.StatusInternalServerError)
		}
		return nil
	}

	return &store{database: newDB}, finallyFn
}

func NewPostgresStore(cfg *config.Config) (repo.DBRepo, func() error) {
	var logLevel glog.LogLevel
	switch cfg.LogLevel {
	case "debug", "DEBUG":
		logLevel = glog.Info
	default:
		logLevel = glog.Warn
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
			cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.DBSSLMode),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: glog.Default.LogMode(logLevel),
	})
	if err != nil {
		zap.L().Panic("cannot connect to db", zap.Error(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Panic("cannot connect to db", zap.Error(err))
	}
	/*
		https://github.com/brettwooldridge/HikariCP/wiki/About-Pool-Sizing
		connections = ((core_count * 2) + effective_spindle_count)
		9 connections = ((4 * 2) + 1)
	*/
	sqlDB.SetMaxIdleConns(9)

	/*
		To avoid the server creating new connections outside of the connection
		pool and creating a "too many clients" deadlock situation, we set max
		open connections to less than the default database limit (100). This
		workload is calculated for up to 4 instances: 3 API servers + 1 indexer
		for up to 72 open connections with 28 reserved.
	*/
	sqlDB.SetMaxOpenConns(18)

	/*
		This is set to avoid zombie connections for inactive, but open connections
		by closing them lazily (will not close active connections). These zombie
		connections may occur from an I/O partition or a single upset event.
	*/
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return &store{database: db}, sqlDB.Close
}
