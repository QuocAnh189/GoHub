package model

import (
	modelUser "gohub/domains/users/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID                 string               `json:"id" gorm:"unique;not null;index;primary_key"`
	ConversationId     string               `json:"conversationId" gorm:"not null"`
	Conversation       *Conversation        `json:"conversation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SenderId           string               `json:"senderId" gorm:"not null"`
	Sender             *modelUser.User      `json:"sender" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId         string               `json:"receiverId" gorm:"not null"`
	Receiver           *modelUser.User      `json:"receiver" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content            string               `json:"content"`
	MessageAttachments []*MessageAttachment `json:"messageAttachments"`
	CreatedAt          time.Time            `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt          time.Time            `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt       `json:"deletedAt" gorm:"index"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (Message) TableName() string {
	return "messages"
}
