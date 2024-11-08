package model

import (
	relation "gohub/domains/shares/model"

	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	ID        		string     							`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     							`json:"name"`
	IsDeleted     	bool       							`json:"isDeleted" gorm:"default:0"`
	Functions       []*relation.CommandInFunction 		`json:"functions" gorm:"many2many:command_functions;"`
	// DeletedAt     	gorm.DeletedAt  				`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  						`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  						`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (Command) TableName() string {
	return "commands"
}