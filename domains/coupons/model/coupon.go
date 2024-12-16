package model

import (
	"time"

	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	ID              string          `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId          string          `json:"userId" gorm:"not null"`
	User            *modelUser.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CoverImageUrl   string          `json:"coverImageUrl" gorm:"not null"`
	Name            string          `json:"name" gorm:"not null"`
	Description     string          `json:"description" gorm:"not null"`
	MinQuantity     int             `json:"minQuantity"`
	MinPrice        float64         `json:"minPrice"`
	PercentageValue float64         `json:"percentageValue"`
	ExpireDate      string          `json:"expireDate" gorm:"not null"`
	CreatedAt       time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt  `json:"deletedAt" gorm:"index"`
}

func (c *Coupon) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()

	return nil
}

func (Coupon) TableName() string {
	return "coupons"
}
