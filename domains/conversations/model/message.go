package model

import (
	modelEvent "gohub/domains/events/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID             string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId        string     			`json:"eventId" gorm:"not null"`
	Event          *modelEvent.Event	`json:"event"`
	ConversationId string     			`json:"conversationId" gorm:"not null"`
	Conversation   *Conversation		`json:"conversation"`
	ContentType    string     			`json:"contentType"`
	ImageUrl       string     			`json:"iconImageUrl"`
	ImageFileName  string     			`json:"iconImageFileName"`
	VideoUrl       string     			`json:"iconVideoUrl"`
	VideoFileName  string     			`json:"iconVideoFileName"`
	AudioUrl       string     			`json:"iconAudioUrl"`
	AudioFileName  string     			`json:"iconAudioFileName"`
	IsDeleted      bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt      *time.Time 			`json:"deletedAt" gorm:"index"`
	// CreatedAt      time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt      time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

func (Message) TableName() string {
	return "messages"
}