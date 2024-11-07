package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	ID        		string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	UserID    		string     			`json:"userId" gorm:"not null"`
	User            *User 				`json:"user"`
	RoleID    		string     			`json:"roleId" gorm:"not null"`
	Role            *Role               `json:"role"`
	IsDeleted    	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ur *UserRole) BeforeCreate(tx * gorm.DB) error {
	ur.ID = uuid.New().String()

	return nil
}

func (UserRole) TableName() string {
	return "user_roles"
}