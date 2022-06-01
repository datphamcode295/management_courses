package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcessedBlock struct {
	ID          string `gorm:"not null" json:"id"`
	BlockNumber int64  `gorm:"not null" json:"block_number"`
}

func (ProcessedBlock) TableName() string {
	return "processed_block"
}

func (o *ProcessedBlock) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
