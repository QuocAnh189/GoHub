package model

import (
	modelEvent "gohub/domains/events/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LabelInEvent struct {
	gorm.Model
	ID        	string     					`json:"id" gorm:"unique;not null;index;primary_key"`
	LabelId    	string     					`json:"labelId" gorm:"not null"`
	Label      	*Label    					`json:"label" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventId    	string     					`json:"eventId" gorm:"not null"`
	Event      	*modelEvent.Event    		`json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (l *LabelInEvent) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()

	return nil
}

func (LabelInEvent) TableName() string {
	return "label_in_events"
}