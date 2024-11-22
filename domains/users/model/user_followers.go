package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFollower struct {
	gorm.Model
	ID         		string    			`json:"id" gorm:"unique;not null;index;primary_key"`
	FollowerId 		string   			`json:"followerId" gorm:"not null"`
	Follower        *User 				`json:"follower" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FolloweeId 		string    			`json:"followeeId" gorm:"not null"`
	Followee        *User 				`json:"followee" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (uf *UserFollower) BeforeCreate(tx * gorm.DB) error {
	uf.ID = uuid.New().String()

	return nil
}

func (UserFollower) TableName() string {
	return "user_followers"
}