package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFavourite struct {
	gorm.Model
	ID         string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId 	   string     		`json:"userId" gorm:"not null"`
	EventId    string     		`json:"eventId" gorm:"not null"`
	IsDeleted  bool       		`json:"isDeleted" gorm:"default:0"`
	// DeletedAt  gorm.DeletedAt 	`json:"deletedAt" gorm:"index"`
	// CreatedAt  time.Time  		`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt  time.Time  		`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ec *EventFavourite) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}


func (EventFavourite) TableName() string {
	return "event_favourites"
}