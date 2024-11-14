package model

import (
	relation "gohub/domains/shares/model"

	"gorm.io/gorm"
)

type Function struct {
	gorm.Model
	ID        		string     							`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     							`json:"name"`
	Url       		string     							`json:"url"`
	SortOrder 		string     							`json:"sortOrder"`
	ParentId  		string     							`json:"parentId"`
	Commands	  	[]*relation.CommandInFunction 		`json:"commands"  gorm:"many2many:command_functions;"`
}

func (Function) TableName() string {
	return "functions"
}