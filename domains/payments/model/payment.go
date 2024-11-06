package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID                  string     `json:"id" gorm:"unique;not null;index;primary_key"`
	TicketQuantity      int        `json:"ticketQuantity" gorm:"not null"`
	CustomerName        string     `json:"customerName" gorm:"not null"`
	CustomerPhone       string     `json:"customerPhone" gorm:"not null"`
	CustomerEmail       string     `json:"customerEmail" gorm:"not null"`
	TotalPrice          float64    `json:"totalPrice" gorm:"not null"`
	DiscountPrice       float64    `json:"discountPrice" gorm:"not null"`
	Status              string     `json:"status" gorm:"default:'pending'"`
	UserPaymentMethodID string     `json:"userPaymentMethodId" gorm:"not null"`
	PaymentSessionID    string     `json:"paymentSessionId" gorm:"not null"`
	EventID             string     `json:"eventId" gorm:"not null"`
	AuthorID            string     `json:"authorId" gorm:"not null"`
	IsDeleted     		bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     		*time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     		time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     		time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}