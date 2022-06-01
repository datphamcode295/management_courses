package pg

import (
	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/repo"
)

type processedBlockRepo struct{}

func NewProcessedBlockRepoo() repo.ProcessedBlockRepo {
	return &processedBlockRepo{}
}

func (pb *processedBlockRepo) Create(s repo.DBRepo, param model.ProcessedBlock) (*model.ProcessedBlock, error) {
	db := s.DB()
	return &param, db.Create(&param).Error
}

func (pb *processedBlockRepo) GetCurrentProcessedBlock(s repo.DBRepo) (*model.ProcessedBlock, error) {
	db := s.DB()
	block := model.ProcessedBlock{}
	return &block, db.First(&block).Error
}

// TODO: optimize this, likely we will have 1s 1 query call situation
func (pb *processedBlockRepo) UpdateBlockNumberByBlockNumber(s repo.DBRepo, blockNumber int64, updatedBlock int64) (*model.ProcessedBlock, error) {
	db := s.DB()
	block := model.ProcessedBlock{}
	return &block, db.Model(&block).Where("block_number = ?", blockNumber).Update("block_number", updatedBlock).Error
}
