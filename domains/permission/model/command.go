package model

import (
	"time"
)

type Command struct {
	ID        		string     `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     `json:"name"`
	IsDeleted     	bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
