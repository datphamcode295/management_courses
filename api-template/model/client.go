package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID       string `gorm:"not null" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Role     string `gorm:"not null" json:"role"`
	ClassID  string `json:"class_id"`
}

func (Client) TableName() string {
	return "client"
}

func (o *Client) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
