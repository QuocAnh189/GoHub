package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID        			string     							`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      			string     							`json:"name" gorm:"not null"`
	IconImageUrl 		string 								`json:"iconImageUrl"`
	IconImageFileName 	string 								`json:"iconImageFileName"`
	Color 				string 								`json:"color"`
	IsDeleted 			bool       							`json:"isDeleted" gorm:"default:0"`
	// DeletedAt 			gorm.DeletedAt 						`json:"deletedAt" gorm:"index"`
	// CreatedAt 			time.Time  							`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt 			time.Time  							`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Category) BeforeCreate() error {
	c.ID = uuid.New().String()
	return nil
}

func (Category) TableName() string {
	return "categories"
}