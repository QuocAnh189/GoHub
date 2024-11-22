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
	PaymentAccountNumber    string    			`json:"paymentAccountNumber" gorm:"not null"`
	PaymentAccountQrCodeUrl string    			`json:"paymentAccountQrCodeUrl" gorm:"not null"`
	CheckoutContent         string    			`json:"checkoutContent" gorm:"not null"`
}

func (up *UserPayment) BeforeCreate(tx * gorm.DB) error {
	up.ID = uuid.New().String()

	return nil
}

func (UserPayment) TableName() string {
	return "user_payments"
}