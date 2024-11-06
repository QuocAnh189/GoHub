package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        			string     	`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      			string     	`json:"name"`
	IconImageUrl 		string 		`json:"iconImageUrl"`
	IconImageFileName 	string 		`json:"iconImageFileName"`
	Color 				string 		`json:"color"`
	EventId 			string 		`json:"eventId"`
	IsDeleted 			bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt 			*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt 			time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 			time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Category) BeforeCreate() error {
	c.ID = uuid.New().String()
	return nil
}