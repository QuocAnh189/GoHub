package model

import (
	"github.com/google/uuid"
	couponModel "gohub/domains/coupons/model"
	"gorm.io/gorm"
	"time"
)

type EventCoupons struct {
	ID        string             `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId   string             `json:"eventId" gorm:"not null"`
	Event     *Event             `json:"event"`
	CouponId  string             `json:"couponId" gorm:"not null"`
	Coupon    couponModel.Coupon `json:"coupon"`
	CreatedAt time.Time          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time          `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt     `json:"deletedAt" gorm:"index"`
}

func (ec *EventCoupons) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()
	return nil
}

func (EventCoupons) TableName() string {
	return "event_coupons"
}
