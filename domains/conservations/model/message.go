package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             string     `json:"id" gorm:"unique;not null;index;primary_key"`
	EventId        string     `json:"eventId"`
	ConversationId string     `json:"conversationId"`
	ContentType    string     `json:"contentType"`
	ImageUrl       string     `json:"iconImageUrl"`
	ImageFileName  string     `json:"iconImageFileName"`
	VideoUrl       string     `json:"iconVideoUrl"`
	VideoFileName  string     `json:"iconVideoFileName"`
	AudioUrl       string     `json:"iconAudioUrl"`
	AudioFileName  string     `json:"iconAudioFileName"`
	IsDeleted      bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt      *time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt      time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (m *Message) BeforeCreate() error {
	m.ID = uuid.New().String()
	return nil
}