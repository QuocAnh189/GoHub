package model

import (
	relation "gohub/domains/shares/model"

	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	ID        		string     							`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     							`json:"name"`
	Functions       []*relation.CommandInFunction 		`json:"functions" gorm:"many2many:command_functions;"`
}

func (Command) TableName() string {
	return "commands"
}