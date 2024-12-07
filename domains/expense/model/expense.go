package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	ID          string         `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId     string         `json:"eventId" gorm:"not null"`
	Title       string         `json:"title" gorm:"not null"`
	Total       float64        `json:"total" gorm:"not null default:0"`
	SubExpenses []*SubExpense  `json:"subExpenses" gorm:"foreignKey:ExpenseId;references:ID"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (ex *Expense) BeforeCreate(db *gorm.DB) (err error) {
	ex.ID = uuid.New().String()
	return nil
}

func (Expense) TableName() string {
	return "expenses"
}
