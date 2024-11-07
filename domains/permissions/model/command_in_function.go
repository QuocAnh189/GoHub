package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommandInFunction struct {
	gorm.Model
	ID 				string 	   			`json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId 		string     			`json:"functionId" gorm:"not null"`
	Function        *Function 			`json:"function"`
	CommandId  		string     			`json:"commandId" gorm:"not null"`
	Command         *Command 			`json:"command"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *CommandInFunction) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()

	return nil
}

func (CommandInFunction) TableName() string {
	return "command_in_functions"
}