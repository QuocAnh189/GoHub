package model

import (
	commandModel "gohub/domains/commands/model"
	functionModel "gohub/domains/functions/model"
	roleModel "gohub/domains/roles/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model	
	ID         		string     						`json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId 		string     						`json:"functionId"`
	Function        *functionModel.Function 		`json:"function" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleId     		string     						`json:"roleId"`
	Role 	  		*roleModel.Role 				`json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CommandId  		string     						`json:"commandId"`
	Command         *commandModel.Command 			`json:"command" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}

func (Permission) TableName() string {
	return "permissions"
}