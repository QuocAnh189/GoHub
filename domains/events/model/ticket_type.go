package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketType struct {
	gorm.Model
	ID                  string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId             string     			`json:"eventId" gorm:"not null"`
	Name                string     			`json:"name"`
	Quantity            int        			`json:"quantity"`
	Price               float64    			`json:"price"`
}

func (t *TicketType) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()

	return nil
}

func (TicketType) TableName() string {
	return "ticket_types"
}