package model

import (
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID             			string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	ConversationId 			string     			`json:"conversationId" gorm:"not null"`
	Conversation   			*Conversation		`json:"conversation"`
	SenderId 				string 				`json:"senderId" gorm:"not null"`
	Sender                  *modelUser.User     `json:"sender"`
	ReceiverId 				string 				`json:"receiverId" gorm:"not null"`
	Receiver                *modelUser.User     `json:"receiver"`
	Content 				string 				`json:"content"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (Message) TableName() string {
	return "messages"
}