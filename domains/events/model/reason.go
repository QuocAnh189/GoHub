package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Reason struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId   string         `json:"eventId" gorm:"not null"`
	Event     *Event         `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content   string         `json:"content" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (r *Reason) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Reason) TableName() string {
	return "reasons"
}
