package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	ID        string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId   string     		`json:"eventId" gorm:"not null"`
	Name      string     		`json:"name" gorm:"not null"`
}

func (r *Reason) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Reason) TableName() string {
	return "reasons"
}