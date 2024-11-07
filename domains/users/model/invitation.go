package model

import (
	modelEvent "gohub/domains/events/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	ID        		string      		`json:"id" gorm:"unique;not null;index;primary_key"`
	InviterId 		string      		`json:"inviterId" gorm:"not null"`
	Inviter         *User               `json:"inviter"`
	InvitedId 		string      		`json:"invitedId" gorm:"not null"`
	Invited         *User               `json:"invited"`
	EventId   		string      		`json:"eventId"`
	Event           *modelEvent.Event   `json:"event"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (i *Invitation) BeforeCreate(tx * gorm.DB) error {
	i.ID = uuid.New().String()

	return nil
}

func (Invitation) TableName() string {
	return "invitations"
}