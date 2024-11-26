package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubExpense struct {
	gorm.Model
	ID          string 		`json:"id" gorm:"unique;not null;index;primary_key"`
	ExpenseId 	string 		`json:"expenseId" gorm:"not null"`	
	Name 		string 		`json:"name" gorm:"not null"`
	Price 		float64 	`json:"price" gorm:"not null"`
}

func (se *SubExpense) BeforeCreate(db *gorm.DB) (err error) {
	se.ID = uuid.New().String()
	return nil
}

func (SubExpense) TableName() string {
    return "sub_expenses"
}