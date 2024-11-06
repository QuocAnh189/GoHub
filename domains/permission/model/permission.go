package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID         string     `json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId string     `json:"function_id"`
	RoleId     string     `json:"role_id"`
	CommandId  string     `json:"command_id"`
	IsDeleted  bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt  time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}