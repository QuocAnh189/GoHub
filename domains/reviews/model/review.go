package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        		string     `json:"id" gorm:"unique;not null;index;primary_key"`
	AuthorId  		string     `json:"authorId"`
	EventId   		string     `json:"eventId"`
	Content   		string     `json:"content"`
	Rate      		float32    `json:"rate" gorm:"default:0"`
	IsDeleted     	bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}