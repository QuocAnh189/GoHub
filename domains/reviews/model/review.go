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
	User          	*modelUser.User     `json:"user"`
	EventId   		string     			`json:"eventId" gorm:"not null"`
	Event         	*modelEvent.Event   `json:"event"`
	Content   		string     			`json:"content"`
	Rate      		float32    			`json:"rate" gorm:"default:0"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Review) TableName() string {
	return "reviews"
}