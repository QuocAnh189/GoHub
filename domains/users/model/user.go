package model

import (
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        		string     `json:"id" gorm:"unique;not null;index;primary_key"`
	Email     		string     `json:"email" gorm:"unique;not null;index:idx_user_email"`
	AvatarUrl 		string     `json:"avatar_url"`
	AvatarFileName  string 	   `json:"avatar_file_name"`
	FullName  		string     `json:"full_name"`
	UserName 		string 	   `json:"user_name"`
	PhoneNumber     string     `json:"phone_number"`
	Dob	   	  		time.Time  `json:"dob"`
	Password  		string     `json:"password"`
	Gender    		int 	   `json:"gender" gorm:"default:0"`
	Status 			int 	   `json:"status" gorm:"default:1"`
	IsDeleted       bool       `json:"is_deleted" gorm:"default:0"`
	DeletedAt       *time.Time  `json:"deleted_at" gorm:"index"`
	CreatedAt       time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}


func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))
	
	return nil
}