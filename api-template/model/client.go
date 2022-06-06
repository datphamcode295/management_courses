package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID      string `gorm:"not null" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	ClassID string `json:"class_id"`
	Class   Class
}

func (Client) TableName() string {
	return "client"
}

func (o *Client) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
