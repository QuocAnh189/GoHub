package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageAttachment struct {
	gorm.Model
	ID 						string 		`json:"id" gorm:"unique;not null;index;primary_key"`
	MessageID               string         `json:"messageId" gorm:"not null"`
	MessageType 			string 		`json:"type" gorm:"not null"`
	MessageUrl 				string 		`json:"messageUrl" gorm:"not null"`
	MessageFileName 		string 		`json:"messageFileName" gorm:"not null"`
}

func (m *MessageAttachment) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (MessageAttachment) TableName() string {
	return "message_attachments"
}