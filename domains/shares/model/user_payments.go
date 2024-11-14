package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPayment struct {
	gorm.Model
	ID                      string    			`json:"id" gorm:"unique;not null;index;primary_key"`
	UserId                	string    			`json:"userId" gorm:"not null"`
	PaymentMethodId         string    			`json:"paymentMethodId" gorm:"not null"`
	PaymentAccountNumber    string    			`json:"paymentAccountNumber"`
	PaymentAccountQrCodeUrl string    			`json:"paymentAccountQrCodeUrl"`
	CheckoutContent         string    			`json:"checkoutContent"`
}

func (up *UserPayment) BeforeCreate(tx * gorm.DB) error {
	up.ID = uuid.New().String()

	return nil
}

func (UserPayment) TableName() string {
	return "user_payments"
}