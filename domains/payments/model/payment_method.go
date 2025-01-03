package model

import (
	relation "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID                 string                  `json:"id" gorm:"unique;not null;index;primary_key"`
	MethodName         string                  `json:"methodName" gorm:"not null"`
	MethodLogoFileName string                  `json:"methodLogoFileName" gorm:"not null"`
	MethodLogoUrl      string                  `json:"methodLogoUrl" gorm:"not null"`
	Users              []*relation.UserPayment `json:"users" gorm:"many2many:user_payments;"`
	CreatedAt          time.Time               `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time               `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt          `json:"deletedAt" gorm:"index"`
}

func (pm *PaymentMethod) BeforeCreate(tx *gorm.DB) error {
	pm.ID = uuid.New().String()

	return nil
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
