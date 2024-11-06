package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPayment struct {
	ID                      string    `json:"id" gorm:"unique;not null;index;primary_key"`
	AuthorId                string    `json:"author_id"`
	PaymentMethodId         string    `json:"payment_method_id"`
	PaymentAccountNumber    string    `json:"payment_account_number"`
	PaymentAccountQrCodeUrl string    `json:"payment_account_qr_code_url"`
	CheckoutContent         string    `json:"checkout_content"`
	IsDeleted               bool      `json:"is_deleted" gorm:"default:0"`
	DeletedAt               *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt               time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (up *UserPayment) BeforeCreate(tx * gorm.DB) error {
	up.ID = uuid.New().String()

	return nil
}