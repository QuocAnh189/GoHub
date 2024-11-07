package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFollower struct {
	gorm.Model
	ID         		string    			`json:"id" gorm:"unique;not null;index;primary_key"`
	FollowerId 		string   			`json:"followerId" gorm:"not null"`
	Follower        *User 				`json:"follower"`
	FollowedId 		string    			`json:"followedId" gorm:"not null"`
	Followed        *User 				`json:"followed"`
	IsDeleted     	bool       			`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  	`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  			`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  			`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}

func (uf *UserFollower) BeforeCreate(tx * gorm.DB) error {
	uf.ID = uuid.New().String()

	return nil
}

func (UserFollower) TableName() string {
	return "user_followers"
}