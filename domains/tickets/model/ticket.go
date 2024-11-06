package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID            string     `json:"id" gorm:"unique;not null;index;primary_key"`
	TicketNo      string     `json:"ticket_no"`
	CustomerName  string     `json:"customer_name"`
	CustomerPhone string     `json:"customer_phone"`
	CustomerEmail string     `json:"customer_email"`
	TicketTypeId  string     `json:"ticket_type_id"`
	EventId       string     `json:"event_id"`
	UserId        string     `json:"user_id"`
	PaymentId     string     `json:"payment_id"`
	Status 		  int 	 	 `json:"status" gorm:"default:1"`
	IsDeleted     bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt     time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()

	return nil
}