package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFavourite struct {
	ID         string     `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId 	   string     `json:"userId"`
	EventId    string     `json:"eventId"`
	IsDeleted  bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt  *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ec *EventFavourite) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}
