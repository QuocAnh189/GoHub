package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        string            `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId    string            `json:"userId" gorm:"not null"`
	User      *modelUser.User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventId   string            `json:"eventId" gorm:"not null"`
	Event     *modelEvent.Event `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content   string            `json:"content" gorm:"not null"`
	Rate      float32           `json:"rate" gorm:"not null"`
	CreatedAt time.Time         `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time         `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt    `json:"deletedAt" gorm:"index"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Review) TableName() string {
	return "reviews"
}
