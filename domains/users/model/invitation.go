package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invitation struct {
	ID        	string     `json:"id" gorm:"unique;not null;index;primary_key"`
	InviterId 	string     `json:"inviter_id"`
	InvitedId 	string     `json:"invited_id"`
	EventId   	string     `json:"event_id"`
	IsDeleted 	bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt   *time.Time  `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (i *Invitation) BeforeCreate(tx * gorm.DB) error {
	i.ID = uuid.New().String()

	return nil
}