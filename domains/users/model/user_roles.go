package model

import (
	"github.com/google/uuid"
	modelRole "gohub/domains/roles/model"
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	UserId    string         `json:"userID" gorm:"not null"`
	User      User           `json:"user"`
	RoleId    string         `json:"roleId" gorm:"not null"`
	Role      modelRole.Role `json:"role"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (ur *UserRole) BeforeCreate(tx *gorm.DB) error {
	ur.ID = uuid.New().String()

	return nil
}

func (UserRole) TableName() string {
	return "user_roles"
}
