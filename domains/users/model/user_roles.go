package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        		string     `json:"id" gorm:"unique;not null;index;primary_key"`
	UserID    		string     `json:"userId"`
	RoleID    		string     `json:"roleId"`
	IsDeleted    	bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ur *UserRole) BeforeCreate(tx * gorm.DB) error {
	ur.ID = uuid.New().String()

	return nil
}