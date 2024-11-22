package model

import (
	modelUser "gohub/domains/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LabelInUser struct {
	gorm.Model
	ID        string     				`json:"id" gorm:"unique;not null;index;primary_key"`
	LabelID   string     				`json:"labelId" gorm:"not null"`
	Label     *Label     				`json:"label" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    string     				`json:"userId" gorm:"not null"`
	User      *modelUser.User     		`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (l *LabelInUser) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()

	return nil
}

func (LabelInUser) TableName() string {
	return "label_in_users"
}