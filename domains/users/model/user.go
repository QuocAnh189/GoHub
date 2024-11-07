package model

import (
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        		string     				`json:"id" gorm:"unique;not null;index;primary_key"`
	Email     		string     				`json:"email" gorm:"unique;not null;index:idx_user_email"`
	AvatarUrl 		string     				`json:"avatarUrl"`
	AvatarFileName  string 	   				`json:"avatarFileName"`
	FullName  		string     				`json:"fullName"`
	UserName 		string 	   				`json:"userName"`
	PhoneNumber     string     				`json:"phoneNumber"`
	Dob	   	  		time.Time  				`json:"dob"`
	Password  		string     				`json:"password"`
	Gender    		int 	   				`json:"gender" gorm:"default:0"`
	Status 			int 	   				`json:"status" gorm:"default:1"`
	UserFollower	[]*UserFollower			`json:"userFollower"`
	UserFollowing	[]*UserFollower			`json:"userFollowing"`
	IsDeleted     	bool       				`json:"isDeleted" gorm:"default:0"`
	// DeletedAt     	gorm.DeletedAt  		`json:"deletedAt" gorm:"index"`
	// CreatedAt     	time.Time  				`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     	time.Time  				`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}


func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))
	
	return nil
}

func (User) TableName() string {
	return "users"
}