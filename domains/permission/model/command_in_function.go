package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommandInFunction struct {
	ID 				string 	   	`json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId 		string     	`json:"functionId"`
	CommandId  		string     	`json:"commandId"`
	IsDeleted     	bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *CommandInFunction) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()

	return nil
}