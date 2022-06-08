package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID       string `gorm:"not null" json:"id"`
	Name     string `json:"name"`
	Lecturer string
}

func (Course) TableName() string {
	return "course"
}

func (o *Course) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
