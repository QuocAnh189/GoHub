package model

import (
	modelEvent "gohub/domains/events/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID             			string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId        			string     			`json:"eventId" gorm:"not null"`
	Event          			*modelEvent.Event	`json:"event"`
	ConversationId 			string     			`json:"conversationId" gorm:"not null"`
	Conversation   			*Conversation		`json:"conversation"`
	Type    	   			string     			`json:"type"`
	MessageUrl 	   			string 				`json:"messageUrl"`
	MessageFileName 	    string 				`json:"messageFileName"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (Message) TableName() string {
	return "messages"
}