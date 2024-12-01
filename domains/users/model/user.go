package model

import (
	modelRole "gohub/domains/roles/model"
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             string            `json:"id" gorm:"unique;not null;index;primary_key"`
	Email          string            `json:"email" gorm:"unique;not null;index:idx_user_email"`
	AvatarUrl      string            `json:"avatarUrl"`
	AvatarFileName string            `json:"avatarFileName"`
	FullName       string            `json:"fullName"`
	UserName       string            `json:"userName" gorm:"unique;not null;index:idx_user_username"`
	PhoneNumber    string            `json:"phoneNumber" gorm:"unique"`
	Dob            *string           `json:"dob"`
	Password       string            `json:"password" gorm:"not null"`
	Gender         *string           `json:"gender" gorm:"default:null"`
	Bio            string            `json:"bio" gorm:""`
	Followers      []*User           `json:"followers" gorm:"many2many:user_followers;joinForeignKey:FolloweeId;joinReferences:FollowerId"`
	Followings     []*User           `json:"followings" gorm:"many2many:user_followers;joinForeignKey:FollowerId;joinReferences:FolloweeId"`
	Roles          []*modelRole.Role `json:"roles" gorm:"many2many:user_roles;"`
	Payments       []*UserPayment    `json:"payments" gorm:"many2many:user_payments;"`
	CreatedAt      time.Time         `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time         `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt    `json:"deletedAt" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (User) TableName() string {
	return "users"
}
