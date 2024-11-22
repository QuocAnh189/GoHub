package model

import (
	relation "gohub/domains/shares/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model	
	ID        			string     					`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      			string     					`json:"name" gorm:"unique;not null"`
	IconImageUrl 		string 						`json:"iconImageUrl" gorm:"not null"`
	IconImageFileName 	string 						`json:"iconImageFileName" gorm:"not null"`
	Color 				string 						`json:"color" gorm:"unique;not null"`
	Event               []*relation.EventCategory 	`json:"events" gorm:"many2many:event_categories;"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

func (Category) TableName() string {
	return "categories"
}