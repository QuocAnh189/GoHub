package model

import (
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID            string     `json:"id" gorm:"unique;not null;index;primary_key"`
	TicketNo      string     `json:"ticketNo"`
	CustomerName  string     `json:"customerName"`
	CustomerPhone string     `json:"customerPhone"`
	CustomerEmail string     `json:"customerEmail"`
	TicketTypeId  string     `json:"ticketTypeId"`
	EventId       string     `json:"eventId"`
	UserId        string     `json:"userId"`
	PaymentId     string     `json:"paymentId"`
	Status 		  int 	 	 `json:"status" gorm:"default:1"`
	IsDeleted     bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	t.TicketNo = utils.GenerateCode("NO")
	return nil
}