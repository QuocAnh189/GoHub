package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID                string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Name              string         `json:"name" gorm:"unique;not null"`
	IconImageUrl      string         `json:"iconImageUrl" gorm:"not null"`
	IconImageFileName string         `json:"iconImageFileName" gorm:"not null"`
	Color             string         `json:"color" gorm:"unique;not null"`
	CreatedAt         time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

func (Category) TableName() string {
	return "categories"
}
