package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//ContractConfig ....
type ContractConfig struct {
	Id        string         `gorm:"PRIMARY_KEY;" json:"id"`
	Name      string         `json:"name"`
	Data      datatypes.JSON `json:"data"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (ContractConfig) TableName() string {
	return "contract_config"
}
func (c *ContractConfig) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return nil
}
