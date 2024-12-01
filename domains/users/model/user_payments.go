package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserPayment struct {
	ID                      string         `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId                  string         `json:"userId" gorm:"not null"`
	PaymentMethodId         string         `json:"paymentMethodId" gorm:"not null"`
	PaymentAccountNumber    string         `json:"paymentAccountNumber" gorm:"not null"`
	PaymentAccountQrCodeUrl string         `json:"paymentAccountQrCodeUrl" gorm:"not null"`
	CheckoutContent         string         `json:"checkoutContent" gorm:"not null"`
	CreatedAt               time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt               time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt               gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (up *UserPayment) BeforeCreate(tx *gorm.DB) error {
	up.ID = uuid.New().String()

	return nil
}

func (UserPayment) TableName() string {
	return "user_payments"
}
