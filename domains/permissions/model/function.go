package model

import (
	"gorm.io/gorm"
)

type Function struct {
	gorm.Model
	ID        		string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     			`json:"name"`
	Url       		string     			`json:"url"`
	SortOrder 		string     			`json:"sortOrder"`
	ParentId  		string     			`json:"parentId"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	Commands	  	[]*Command 			`json:"commands"  gorm:"many2many:command_in_functions;"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (Function) TableName() string {
	return "functions"
}