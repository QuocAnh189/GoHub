package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	ID     string `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId string `json:"userID" gorm:"not null"`
	RoleId string `json:"roleId" gorm:"not null"`
}

func (ur *UserRole) BeforeCreate(tx *gorm.DB) error {
	ur.ID = uuid.New().String()

	return nil
}

func (UserRole) TableName() string {
	return "user_roles"
}