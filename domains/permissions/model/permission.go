package model

import (
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID         		string     			`json:"id" gorm:"unique;not null;index;primary_key"`
	FunctionId 		string     			`json:"functionId" gorm:"not null"`
	Function        *Function 			`json:"function"`
	RoleId     		string     			`json:"roleId" gorm:"not null"`
	Role 	  		*modelUser.Role 	`json:"role"`
	CommandId  		string     			`json:"commandId" gorm:"not null"`
	Command         *Command 			`json:"command"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()

	return nil
}

func (Permission) TableName() string {
	return "permissions"
}