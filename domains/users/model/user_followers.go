package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFollower struct {
	ID         string    `json:"id" gorm:"unique;not null;index;primary_key"`
	FollowerId string    `json:"follower_id"`
	FollowedId string    `json:"followed_id"`
	IsDeleted  bool      `json:"is_deleted" gorm:"default:0"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"index"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (uf *UserFollower) BeforeCreate(tx * gorm.DB) error {
	uf.ID = uuid.New().String()

	return nil
}