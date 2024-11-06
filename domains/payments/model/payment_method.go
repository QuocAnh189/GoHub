package model

import "time"

type PaymentMethod struct {
	ID                 string     `json:"id" gorm:"unique;not null;index;primary_key"`
	MethodName         string     `json:"methodName" gorm:"not null"`
	MethodLogoFileName string     `json:"methodLogoFileName" gorm:"not null"`
	MethodLogoUrl      string     `json:"methodLogoUrl" gorm:"not null"`
	IsDeleted    	   bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt          *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt          time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}