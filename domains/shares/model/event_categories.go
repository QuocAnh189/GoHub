package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventCategory struct {
	gorm.Model
	ID         string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	CategoryId string     		`json:"categoryId" gorm:"not null"`
	EventId    string     		`json:"eventId" gorm:"not null"`
}

func (ec *EventCategory) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}

func (EventCategory) TableName() string {
	return "event_categories"
}