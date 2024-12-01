package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MessageAttachment struct {
	ID              string         `json:"id" gorm:"unique;not null;index;primary_key"`
	MessageID       string         `json:"messageId" gorm:"not null"`
	Message         *Message       `json:"message" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MessageType     string         `json:"type" gorm:"not null"`
	MessageUrl      string         `json:"messageUrl" gorm:"not null"`
	MessageFileName string         `json:"messageFileName" gorm:"not null"`
	CreatedAt       time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (m *MessageAttachment) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (MessageAttachment) TableName() string {
	return "message_attachments"
}
