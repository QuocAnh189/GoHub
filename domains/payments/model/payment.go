package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID                  string            `json:"id" gorm:"unique;not null;index;primary_key"`
	EventID             string            `json:"eventId" gorm:"not null"`
	Event               *modelEvent.Event `json:"event"`
	CustomerName        string            `json:"customerName" gorm:"not null"`
	CustomerEmail       string            `json:"customerEmail" gorm:"not null"`
	CustomerPhone       string            `json:"customerPhone" gorm:"not null"`
	UserId              string            `json:"userId" gorm:"not null"`
	User                *modelUser.User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentSessionID    string            `json:"paymentSessionId" gorm:"not null"`
	UserPaymentMethodID string            `json:"userPaymentMethodId" gorm:"not null"`
	UserPaymentMethod   *PaymentMethod    `json:"userPaymentMethod" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TicketQuantity      int               `json:"ticketQuantity" gorm:"not null"`
	TotalPrice          float64           `json:"totalPrice" gorm:"not null"`
	DiscountPrice       float64           `json:"discountPrice" gorm:"not null"`
	Status              string            `json:"status" gorm:"default:'PENDING'"`
	CreatedAt           time.Time         `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt           time.Time         `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt    `json:"deletedAt" gorm:"index"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}

func (Payment) TableName() string {
	return "payments"
}
