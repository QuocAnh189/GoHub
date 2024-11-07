package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LabelInEvent struct {
	gorm.Model
	ID        	string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	LabelId    	string     		`json:"labelId" gorm:"not null"`
	EventId    	string     		`json:"eventId" gorm:"not null"`
	IsDeleted 	bool       		`json:"isDeleted" gorm:"default:0"`
	// DeletedAt 	gorm.DeletedAt  `json:"deletedAt" gorm:"index"`
	// CreatedAt 	time.Time  		`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt 	time.Time  		`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (l *LabelInEvent) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()

	return nil
}

func (LabelInEvent) TableName() string {
	return "label_in_events"
}