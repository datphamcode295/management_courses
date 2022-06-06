package model

import (
	"gorm.io/gorm"
)

type Class struct {
	ID string
}

func (Class) TableName() string {
	return "class"
}

func (o *Class) BeforeCreate(tx *gorm.DB) error {
	// o.ID = uuid.New().String()
	return nil
}
