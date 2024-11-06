package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
	AuthorId  string     `json:"author_id"`
	EventId   string     `json:"event_id"`
	Content   string     `json:"content"`
	Rate      float32    `json:"rate" gorm:"default:0"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}