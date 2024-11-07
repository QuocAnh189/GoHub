package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	ID                 string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	MethodName         string     		`json:"methodName" gorm:"not null"`
	MethodLogoFileName string     		`json:"methodLogoFileName" gorm:"not null"`
	MethodLogoUrl      string     		`json:"methodLogoUrl" gorm:"not null"`
	IsDeleted    	   bool       		`json:"isDeleted" gorm:"default:0"`
	// DeletedAt          gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt          time.Time  		`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt          time.Time  		`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (pm *PaymentMethod) BeforeCreate(tx *gorm.DB) error {
	pm.ID = uuid.New().String()

	return nil
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}