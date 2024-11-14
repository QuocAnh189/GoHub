package model

import (
	relation "gohub/domains/shares/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	ID                 	string     					`json:"id" gorm:"unique;not null;index;primary_key"`
	MethodName         	string     					`json:"methodName" gorm:"not null"`
	MethodLogoFileName 	string     					`json:"methodLogoFileName" gorm:"not null"`
	MethodLogoUrl      	string     					`json:"methodLogoUrl" gorm:"not null"`
	Users               []*relation.UserPayment		`json:"users" gorm:"many2many:user_payments;"`
}

func (pm *PaymentMethod) BeforeCreate(tx *gorm.DB) error {
	pm.ID = uuid.New().String()

	return nil
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}