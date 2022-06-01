package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          string    `gorm:"not null" json:"id"`
	BlockNumber uint64    `json:"block_number"`
	TxHash      string    `json:"tx_hash"`
	ActionName  string    `json:"action_name"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func (o *Transaction) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
