package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TicketType struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId   string         `json:"eventId" gorm:"not null"`
	Event     *Event         `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string         `json:"name" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Sale      int            `json:"sale" gorm:"not null default:0"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (t *TicketType) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()

	return nil
}

func (TicketType) TableName() string {
	return "ticket_types"
}
