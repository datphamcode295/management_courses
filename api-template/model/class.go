package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID   string `gorm:"PRIMARY_KEY;" json:"Id"`
	Name string
}

func (Class) TableName() string {
	return "class"
}

func (c *Class) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}
