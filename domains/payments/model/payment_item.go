package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentItem struct {
	ID            string     `json:"id" gorm:"unique;not null;index;primary_key"`
	PaymentID     string     `json:"paymentId" gorm:"not null"`
	EventID       string     `json:"eventId" gorm:"not null"`
	UserID        string     `json:"userId" gorm:"not null"`
	TicketTypeID  string     `json:"ticketTypeId" gorm:"not null"`
	Name          string     `json:"name" gorm:"not null"`
	Quantity      int        `json:"quantity" gorm:"not null"`
	TotalPrice    float64    `json:"total"`
	DiscountPrice float64    `json:"discount" gorm:"not null"`
	IsDeleted     bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (pi *PaymentItem) BeforeCreate(tx *gorm.DB) error {
	pi.ID = uuid.New().String()

	return nil
}