package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	ID      string `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId string `json:"eventId" gorm:"not null"`
	Event   *Event `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string `json:"content" gorm:"not null"`
}

func (r *Reason) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Reason) TableName() string {
	return "reasons"
}
