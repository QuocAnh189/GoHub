package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Label struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      string     `json:"name"`
	IsDeleted bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (l *Label) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()

	return nil
}