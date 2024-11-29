package model

import (
	"time"

	relation "gohub/domains/shares/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	ID              string                   `json:"id" gorm:"unique;not null;index;primary_key"`
	Name            string                   `json:"name" gorm:"not null"`
	Description     string                   `json:"description" gorm:"not null"`
	UserId          string                   `json:"userId" gorm:"not null"`
	User            *modelUser.User          `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MinQuantity     int                      `json:"minQuantity"`
	MinValue        float64                  `json:"minValue"`
	PercentageValue float64                  `json:"percentValue"`
	RealValue       float64                  `json:"realValue"`
	ExpireDate      time.Time                `json:"expireDate" gorm:"not null"`
	Events          []*relation.EventCoupons `json:"events" gorm:"many2many:event_coupons;"`
}

func (c *Coupon) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()

	return nil
}

func (Coupon) TableName() string {
	return "coupons"
}
