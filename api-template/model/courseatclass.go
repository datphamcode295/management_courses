package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseAtClass struct {
	ID       string `gorm:"not null" json:"id"`
	ClassId  string
	CourseId string
}

func (CourseAtClass) TableName() string {
	return "course_at_class"
}

func (o *CourseAtClass) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
