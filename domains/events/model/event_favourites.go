package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EventFavourite struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId    string         `json:"userId" gorm:"not null"`
	EventId   string         `json:"eventId" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (ec *EventFavourite) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}

func (EventFavourite) TableName() string {
	return "event_favourites"
}
