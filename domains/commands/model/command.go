package model

import (
	relation "gohub/domains/shares/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	ID        		string     							`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     							`json:"name" gorm:"not null"`
	Functions       []*relation.CommandInFunction 		`json:"functions" gorm:"many2many:command_functions;"`
}

func (c *Command) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

func (Command) TableName() string {
	return "commands"
}