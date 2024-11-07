package model

import (
	modelEvent "gohub/domains/events/model"
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	ID            string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId       string     		`json:"eventId" gorm:"not null"`
	Event         *modelEvent.Event `json:"event"`
	UserId        string     		`json:"userId" gorm:"not null"`
	User          *modelUser.User  	`json:"user"`
	HostId        string     		`json:"hostId" gorm:"not null"`
	Host          *modelUser.User  	`json:"host"`
	LastMessageId string     		`json:"lastMessageId"`
	LastMessage   *Message			`json:"lastMessage"`   		
	IsDeleted     bool       		`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     gorm.DeletedAt 	`json:"deletedAt" gorm:"index"`
	// CreatedAt     time.Time  		`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     time.Time  		`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Conversation) BeforeCreate() error {
	c.ID = uuid.New().String()
	return nil
}

func (Conversation) TableName() string {
	return "conversations"
}