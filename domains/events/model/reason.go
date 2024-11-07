package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	ID        string     		`json:"id" gorm:"unique;not null;index;primary_key"`
	EventId   string     		`json:"eventId" gorm:"not null"`
	Name      string     		`json:"name"`
	IsDeleted bool       		`json:"isDeleted" gorm:"default:0"`
	// DeletedAt gorm.DeletedAt 	`json:"deletedAt" gorm:"index"`
	// CreatedAt time.Time  		`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt time.Time  		`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Reason) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Reason) TableName() string {
	return "reasons"
}