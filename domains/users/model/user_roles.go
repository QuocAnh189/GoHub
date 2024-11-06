package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        string    `json:"id" gorm:"unique;not null;index;primary_key"`
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:0"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (ur *UserRole) BeforeCreate(tx * gorm.DB) error {
	ur.ID = uuid.New().String()

	return nil
}