package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type transactionRepo struct{}

func NewTransactionRepo() repo.TransactionRepo {
	return &transactionRepo{}
}

func (t *transactionRepo) Create(s repo.DBRepo, param *model.Transaction) (*model.Transaction, error) {
	db := s.DB()
	return param, db.Create(&param).Error
}
