package model

import (
	modelEvent "gohub/domains/events/model"
	modelPayment "gohub/domains/payments/model"
	modelUser "gohub/domains/users/model"
	"gohub/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID            		string     				`json:"id" gorm:"unique;not null;index;primary_key"`
	TicketNo      		string     				`json:"ticketNo"`
	CustomerName  		string     				`json:"customerName"`
	CustomerPhone 		string     				`json:"customerPhone"`
	CustomerEmail 		string     				`json:"customerEmail"`
	TicketTypeId  		string     				`json:"ticketTypeId" gorm:"not null"`
	TicketType    		*modelEvent.TicketType  `json:"ticketType"`
	EventId       		string     				`json:"eventId" gorm:"not null"`
	Event         		*modelEvent.Event      	`json:"event"`
	UserId        		string     				`json:"userId" gorm:"not null"`
	User          		*modelUser.User        	`json:"user"`
	PaymentId     		string     				`json:"paymentId" gorm:"not null"`
	Payment       		*modelPayment.Payment 	`json:"payment"`
	Status 		  		string 	 	 			`json:"status"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	t.TicketNo = utils.GenerateCode("NO")
	return nil
}

func (Ticket) TableName() string {
	return "tickets"
}