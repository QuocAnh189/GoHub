package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentItem struct {
	gorm.Model
	ID            string     				`json:"id" gorm:"unique;not null;index;primary_key"`
	PaymentID     string     				`json:"paymentId" gorm:"not null"`
	Payment       *Payment  				`json:"payment"`
	EventID       string     				`json:"eventId" gorm:"not null"`
	Event         *modelEvent.Event			`json:"event"`
	UserID        string     				`json:"userId" gorm:"not null"`
	User          *modelUser.User			`json:"user"`
	TicketTypeID  string     				`json:"ticketTypeId" gorm:"not null"`
	TicketType    *modelEvent.TicketType    `json:"ticketType"`
	Name          string     				`json:"name"`
	Quantity      int        				`json:"quantity"`
	TotalPrice    float64    				`json:"total"`
	DiscountPrice float64    				`json:"discount"`
	IsDeleted     bool       				`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     gorm.DeletedAt  			`json:"deletedAt" gorm:"index"`
	// CreatedAt     time.Time  				`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     time.Time  				`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (pi *PaymentItem) BeforeCreate(tx *gorm.DB) error {
	pi.ID = uuid.New().String()

	return nil
}

func (PaymentItem) TableName() string {
	return "payment_items"
}