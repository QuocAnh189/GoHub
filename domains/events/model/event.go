package model

import (
	modelCategory "gohub/domains/categories/model"
	modelCoupon "gohub/domains/coupons/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID                 string                    `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId             string                    `json:"userId" gorm:"not null"`
	User               *modelUser.User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name               string                    `json:"name" gorm:"not null"`
	Description        string                    `json:"description"`
	CoverImageUrl      string                    `json:"coverImageUrl" gorm:"not null"`
	CoverImageFileName string                    `json:"coverImageFileName" gorm:"not null"`
	StartTime          *time.Time                `json:"startDate" gorm:"not null"`
	EndTime            *time.Time                `json:"endDate" gorm:"not null"`
	Location           string                    `json:"location" gorm:"not null"`
	PathLocation       string                    `json:"pathLocation" gorm:"not null"`
	EventCycleType     string                    `json:"eventCycleType" gorm:"not null"`
	EventPaymentType   string                    `json:"eventPaymentType" gorm:"not null"`
	IsPrivate          bool                      `json:"isPrivate" gorm:"default:0"`
	EventSubImages     []*EventSubImage          `json:"eventSubImages"`
	Categories         []*modelCategory.Category `json:"categories" gorm:"many2many:event_categories;"`
	Reasons            []*Reason                 `json:"reasons"`
	Coupons            []*modelCoupon.Coupon     `json:"coupons" gorm:"many2many:event_coupons;"`
	Expenses           []*EventExpense           `json:"expenses"`
	TicketTypes        []*TicketType             `json:"ticketTypes"`
	UserFavourite      []*modelUser.User         `json:"userFavourite" gorm:"many2many:event_favourites;"`
	CreatedAt          time.Time                 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time                 `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt            `json:"deletedAt" gorm:"index"`
}

func (e *Event) BeforeCreate(db *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return nil
}

func (Event) TableName() string {
	return "events"
}
