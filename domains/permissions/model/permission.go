package model

import (
	commandModel "gohub/domains/commands/model"
	functionModel "gohub/domains/functions/model"
	roleModel "gohub/domains/roles/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID         string                  `json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId string                  `json:"functionId"`
	Function   *functionModel.Function `json:"function" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleId     string                  `json:"roleId"`
	Role       *roleModel.Role         `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommandId  string                  `json:"commandId"`
	Command    *commandModel.Command   `json:"command" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt  time.Time               `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time               `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt          `json:"deletedAt" gorm:"index"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}

func (Permission) TableName() string {
	return "permissions"
}
