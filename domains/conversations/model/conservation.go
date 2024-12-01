package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID            string            `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId       string            `json:"eventId"`
	Event         *modelEvent.Event `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId        string            `json:"userId" gorm:"not null"`
	User          *modelUser.User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizerId   string            `json:"organizerId" gorm:"not null"`
	Organizer     *modelUser.User   `json:"organizer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LastMessageId string            `json:"lastMessageId"`
	LastMessage   *Message          `json:"lastMessage" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time         `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time         `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt    `json:"deletedAt" gorm:"index"`
}

func (c *Conversation) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

func (Conversation) TableName() string {
	return "conversations"
}
