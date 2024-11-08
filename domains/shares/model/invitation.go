package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	ID        		string      		`json:"id" gorm:"unique;not null;index;primary_key"`
	InviterId 		string      		`json:"inviterId" gorm:"not null"`
	InvitedId 		string      		`json:"invitedId" gorm:"not null"`
	EventId   		string      		`json:"eventId"`
}

func (i *Invitation) BeforeCreate(tx * gorm.DB) error {
	i.ID = uuid.New().String()

	return nil
}

func (Invitation) TableName() string {
	return "invitations"
}