package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID        		string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId  		string     			`json:"userId" gorm:"not null"`
	User          	*modelUser.User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventId   		string     			`json:"eventId" gorm:"not null"`
	Event         	*modelEvent.Event   `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content   		string     			`json:"content" gorm:"not null"`
	Rate      		float32    			`json:"rate" gorm:"not null"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Review) TableName() string {
	return "reviews"
}