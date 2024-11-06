package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventSubImage struct {
	ID         		string     	`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId    		string     	`json:"eventId"`
	ImageUrl   		string     	`json:"imageUrl"`
	ImageFileName 	string 		`json:"imageFileName"`
	IsDeleted  		bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt  		*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt  		time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  		time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ec *EventSubImage) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()

	return nil
}
