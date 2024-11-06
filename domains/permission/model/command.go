package model

import (
	"time"
)

type Command struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      string     `json:"name"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
