package model

import (
	"github.com/google/uuid"
	modelUser "gohub/domains/users/model"
	"gorm.io/gorm"
	"time"
)

type Invitation struct {
	ID        string          `json:"id" gorm:"unique;not null;index;primary_key"`
	InviterId string          `json:"inviterId" gorm:"not null"`
	Inviter   *modelUser.User `json:"inviter" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InviteeId string          `json:"inviteeId" gorm:"not null"`
	Invitee   *modelUser.User `json:"invitee" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventId   string          `json:"eventId" gorm:"not null"`
	Event     *Event          `json:"event" gorm:"constraint:OnUpdate:CASCADE;"`
	CreatedAt time.Time       `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `json:"deletedAt" gorm:"index"`
}

func (i *Invitation) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.New().String()

	return nil
}

func (Invitation) TableName() string {
	return "invitations"
}
