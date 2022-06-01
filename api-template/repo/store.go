package repo

import (
	"context"

	"gorm.io/gorm"
)

type DBRepo interface {
	DB() *gorm.DB
	NewTransaction() (DBRepo, FinallyFunc)
	NewTransactionWithContext(ctx context.Context) (DBRepo, FinallyFunc)
}

type FinallyFunc = func(error) error
