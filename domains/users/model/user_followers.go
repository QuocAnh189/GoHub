package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserFollower struct {
	ID         string         `json:"id" gorm:"unique;not null;index;primary_key"`
	FollowerId string         `json:"followerId" gorm:"not null"`
	Follower   *User          `json:"follower" gorm:"foreignKey:FollowerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FolloweeId string         `json:"followeeId" gorm:"not null"`
	Followee   *User          `json:"followee" gorm:"foreignKey:FolloweeId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (uf *UserFollower) BeforeCreate(tx *gorm.DB) error {
	uf.ID = uuid.New().String()

	return nil
}

func (UserFollower) TableName() string {
	return "user_followers"
}
