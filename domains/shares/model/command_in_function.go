package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommandInFunction struct {
	gorm.Model
	ID         string `json:"id" gorm:"unique;not null;index;primary_key"`
	CommandID  string `json:"command_id" gorm:"not null"`
	FunctionID string `json:"function_id" gorm:"not null"`
}

func (cf *CommandInFunction) BeforeCreate(tx *gorm.DB) error {
	cf.ID = uuid.New().String()

	return nil
}

func (CommandInFunction) TableName() string {
	return "command_functions"
}