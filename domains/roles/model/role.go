package model

import (
	relation "gohub/domains/shares/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID        string               `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      string               `json:"name"`
	User      []*relation.UserRole `json:"users" gorm:"many2many:user_roles;"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}

func (Role) TableName() string {
	return "roles"
}
