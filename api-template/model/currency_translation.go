package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//CurrencyTranslation ....
type CurrencyTranslation struct {
	Id              string    `gorm:"PRIMARY_KEY;" json:"id"`
	Name            string    `json:"name"`
	Currency        string    `json:"currency"`
	CurrencyAddress string    `json:"currency_address"`
	CurrencyImage   string    `json:"currency_image"`
	PriceUsd        float64   `json:"price_usd"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (CurrencyTranslation) TableName() string {
	return "currency_translation"
}
func (c *CurrencyTranslation) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return nil
}
