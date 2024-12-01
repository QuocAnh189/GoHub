package model

import (
	"github.com/google/uuid"
	modelFunction "gohub/domains/functions/model"
	"gorm.io/gorm"
	"time"
)

type CommandInFunction struct {
	ID         string                  `json:"id" gorm:"unique;not null;index;primary_key"`
	CommandID  string                  `json:"command_id" gorm:"not null"`
	Command    *Command                `json:"command"`
	FunctionID string                  `json:"function_id" gorm:"not null"`
	Function   *modelFunction.Function `json:"function"`
	CreatedAt  time.Time               `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time               `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt          `json:"deletedAt" gorm:"index"`
}

func (cf *CommandInFunction) BeforeCreate(tx *gorm.DB) error {
	cf.ID = uuid.New().String()

	return nil
}

func (CommandInFunction) TableName() string {
	return "command_functions"
}
