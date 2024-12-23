package model

import (
	modelEvent "gohub/domains/events/model"
	modelPayment "gohub/domains/payments/model"
	modelUser "gohub/domains/users/model"
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID            string                 `json:"id" gorm:"unique;not null;index;primary_key"`
	TicketNo      string                 `json:"ticketNo" gorm:"not null"`
	CustomerName  string                 `json:"customerName" gorm:"not null"`
	CustomerPhone string                 `json:"customerPhone" gorm:"not null"`
	CustomerEmail string                 `json:"customerEmail" gorm:"not null"`
	TicketTypeId  string                 `json:"ticketTypeId" gorm:"not null"`
	TicketType    *modelEvent.TicketType `json:"ticketType" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventId       string                 `json:"eventId" gorm:"not null"`
	Event         *modelEvent.Event      `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId        string                 `json:"userId" gorm:"not null"`
	User          *modelUser.User        `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaymentId     string                 `json:"paymentId" gorm:"not null"`
	Payment       *modelPayment.Payment  `json:"payment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time              `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time              `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt         `json:"deletedAt" gorm:"index"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	t.TicketNo = utils.GenerateCode("NO")
	return nil
}

func (Ticket) TableName() string {
	return "tickets"
}
