package model

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID            string     `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId       string     `json:"eventId"`
	UserId        string     `json:"userId"`
	HostId        string     `json:"hostId"`
	LastMessageId string     `json:"lastMessageId"`
	IsDeleted     bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Conversation) BeforeCreate() error {
	c.ID = uuid.New().String()
	return nil
}