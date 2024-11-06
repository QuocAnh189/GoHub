package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPayment struct {
	ID                      string    	`json:"id" gorm:"unique;not null;index;primary_key"`
	AuthorId                string    	`json:"authorId"`
	PaymentMethodId         string    	`json:"paymentMethodId"`
	PaymentAccountNumber    string    	`json:"paymentAccountNumber"`
	PaymentAccountQrCodeUrl string    	`json:"paymentAccountQrCodeUrl"`
	CheckoutContent         string    	`json:"checkoutContent"`
	IsDeleted     			bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt     			*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt     			time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     			time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (up *UserPayment) BeforeCreate(tx * gorm.DB) error {
	up.ID = uuid.New().String()

	return nil
}