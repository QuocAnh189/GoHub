package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentLine struct {
	gorm.Model
	ID            	string     					`json:"id" gorm:"unique;not null;index;primary_key"`
	PaymentID     	string     					`json:"paymentId" gorm:"not null"`
	Payment       	*Payment  					`json:"payment"`
	EventID       	string     					`json:"eventId" gorm:"not null"`
	Event         	*modelEvent.Event			`json:"event"`
	UserID        	string     					`json:"userId" gorm:"not null"`
	User          	*modelUser.User				`json:"user"`
	TicketTypeID  	string     					`json:"ticketTypeId" gorm:"not null"`
	TicketType    	*modelEvent.TicketType    	`json:"ticketType"`
	Quantity      	int        					`json:"quantity"`
	TotalPrice    	float64    					`json:"total"`
}

func (pi *PaymentLine) BeforeCreate(tx *gorm.DB) error {
	pi.ID = uuid.New().String()

	return nil
}

func (PaymentLine) TableName() string {
	return "payment_lines"
}