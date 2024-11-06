package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFollower struct {
	ID         		string    	`json:"id" gorm:"unique;not null;index;primary_key"`
	FollowerId 		string   	`json:"followerId"`
	FollowedId 		string    	`json:"followedId"`
	IsDeleted     	bool       	`json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time 	`json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  	`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  	`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (uf *UserFollower) BeforeCreate(tx * gorm.DB) error {
	uf.ID = uuid.New().String()

	return nil
}