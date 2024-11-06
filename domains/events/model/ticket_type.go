package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketType struct {
	ID                  string     `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId             string     `json:"eventId"`
	Name                string     `json:"name"`
	Quantity            int        `json:"quantity"`
	Price               float64    `json:"price"`
	NumberOfSoldTickets int        `json:"numberOfSoldTickets"`
	IsDeleted           bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt           *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt           time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *TicketType) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()

	return nil
}