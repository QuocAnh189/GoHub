package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentLine struct {
	ID           string                 `json:"id" gorm:"unique;not null;index;primary_key"`
	PaymentID    string                 `json:"paymentId" gorm:"not null"`
	Payment      *Payment               `json:"payment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventID      string                 `json:"eventId" gorm:"not null"`
	Event        *modelEvent.Event      `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID       string                 `json:"userId" gorm:"not null"`
	User         *modelUser.User        `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TicketTypeID string                 `json:"ticketTypeId" gorm:"not null"`
	TicketType   *modelEvent.TicketType `json:"ticketType" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity     int                    `json:"quantity"`
	TotalPrice   float64                `json:"total"`
	CreatedAt    time.Time              `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time              `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt         `json:"deletedAt" gorm:"index"`
}

func (pi *PaymentLine) BeforeCreate(tx *gorm.DB) error {
	pi.ID = uuid.New().String()

	return nil
}

func (PaymentLine) TableName() string {
	return "payment_lines"
}
