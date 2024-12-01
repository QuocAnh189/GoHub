package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SubExpense struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	ExpenseId string         `json:"expenseId" gorm:"not null"`
	Name      string         `json:"name" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (se *SubExpense) BeforeCreate(db *gorm.DB) (err error) {
	se.ID = uuid.New().String()
	return nil
}

func (SubExpense) TableName() string {
	return "sub_expenses"
}
