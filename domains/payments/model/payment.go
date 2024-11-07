package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID                  string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	TicketQuantity      int        			`json:"ticketQuantity" gorm:"not null"`
	CustomerName        string     			`json:"customerName" gorm:"not null"`
	CustomerPhone       string     			`json:"customerPhone" gorm:"not null"`
	CustomerEmail       string     			`json:"customerEmail" gorm:"not null"`
	TotalPrice          float64    			`json:"totalPrice" gorm:"not null"`
	DiscountPrice       float64    			`json:"discountPrice" gorm:"not null"`
	Status              string     			`json:"status" gorm:"default:'pending'"`
	UserPaymentMethodID string     			`json:"userPaymentMethodId" gorm:"not null"`
	UserPaymentMethod   *PaymentMethod		`json:"userPaymentMethod"`
	PaymentSessionID    string     			`json:"paymentSessionId" gorm:"not null"`
	EventID             string     			`json:"eventId" gorm:"not null"`
	Event               *modelEvent.Event	`json:"event"`
	UserId            	string     			`json:"userId" gorm:"not null"`
	User                *modelUser.User		`json:"user"`
	IsDeleted     		bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     		gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     		time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     		time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}

func (Payment) TableName() string {
	return "payments"
}