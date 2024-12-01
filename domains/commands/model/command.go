package model

import (
	functionModel "gohub/domains/functions/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Command struct {
	ID        string                    `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      string                    `json:"name" gorm:"not null"`
	Functions []*functionModel.Function `json:"functions" gorm:"many2many:command_functions;"`
	CreatedAt time.Time                 `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time                 `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt            `json:"deletedAt" gorm:"index"`
}

func (c *Command) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

func (Command) TableName() string {
	return "commands"
}
