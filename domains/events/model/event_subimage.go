package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventSubImage struct {
	gorm.Model
	ID         		string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId    		string     			`json:"eventId" gorm:"not null"`
	Event       	*Event     			`json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ImageUrl   		string     			`json:"imageUrl" gorm:"not null"`
	ImageFileName 	string 				`json:"imageFileName" gorm:"not null"`
}

func (ec *EventSubImage) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}

func (EventSubImage) TableName() string {
	return "event_subimages"
}