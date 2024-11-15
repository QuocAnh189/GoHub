package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventCoupons struct {
	gorm.Model
	ID       string `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId  string `json:"eventId"`
	CouponId string `json:"couponId"`
}

func (ec *EventCoupons) BeforeCreate(tx *gorm.DB) error {
	ec.ID = uuid.New().String()
	return nil;
}

func (EventCoupons) TableName() string {
	return "event_coupons"
}