package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LabelInUser struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
	LabelID   string     `json:"labelId"`
	UserID    string     `json:"userId"`
	IsDeleted bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (l *LabelInUser) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()

	return nil
}