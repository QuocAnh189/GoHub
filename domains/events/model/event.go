package model

import (
	modelCategory "gohub/domains/categories/model"
	modelCoupon "gohub/domains/coupons/model"
	modelExpense "gohub/domains/expense/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID                 string                    `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId             string                    `json:"userId" gorm:"not null"`
	User               *modelUser.User           `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name               string                    `json:"name" gorm:"not null"`
	Description        string                    `json:"description"`
	CoverImageUrl      string                    `json:"coverImageUrl" gorm:"not null"`
	CoverImageFileName string                    `json:"coverImageFileName" gorm:"not null"`
	StartTime          string                    `json:"startTime" gorm:"not null"`
	EndTime            string                    `json:"endTime" gorm:"not null"`
	Location           string                    `json:"location" gorm:"not null"`
	PathLocation       string                    `json:"pathLocation" gorm:"not null"`
	EventCycleType     string                    `json:"eventCycleType" gorm:"not null"`
	EventPaymentType   string                    `json:"eventPaymentType" gorm:"not null"`
	IsPrivate          bool                      `json:"isPrivate" gorm:"default:0"`
	SubImages          []*EventSubImage          `json:"subImages"`
	Categories         []*modelCategory.Category `json:"categories" gorm:"many2many:event_categories;"`
	Reasons            []*Reason                 `json:"reasons"`
	Coupons            []*modelCoupon.Coupon     `json:"coupons" gorm:"many2many:event_coupons;"`
	Expenses           []*modelExpense.Expense   `json:"expenses"`
	TicketTypes        []*TicketType             `json:"ticketTypes"`
	Reviews            []*Review                 `json:"reviews"`
	UserFavourite      []*modelUser.User         `json:"userFavourite" gorm:"many2many:event_favourites;"`
	AverageRate        float64                   `json:"averageRate" sql:"average_rate"`
	CreatedAt          time.Time                 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time                 `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt            `json:"deletedAt" gorm:"index"`
}

type Review struct {
	ID         string         `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId     string         `json:"userId" gorm:"not null"`
	EventId    string         `json:"eventId" gorm:"not null"`
	Content    string         `json:"content" gorm:"not null"`
	Rate       float32        `json:"rate" gorm:"not null"`
	IsPositive bool           `json:"isPositive" gorm:"not null"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (e *Event) BeforeCreate(db *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return nil
}

func (Event) TableName() string {
	return "events"
}
