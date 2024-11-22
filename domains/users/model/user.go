package model

import (
	relation "gohub/domains/shares/model"
	"gohub/pkg/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        		 string     						`json:"id" gorm:"unique;not null;index;primary_key"`
	Email     		 string     						`json:"email" gorm:"unique;not null;index:idx_user_email"`
	AvatarUrl 		 string     						`json:"avatarUrl"`
	AvatarFileName   string 	   						`json:"avatarFileName"`
	FullName  		 string     						`json:"fullName"`
	UserName 		 string 	   						`json:"userName" gorm:"unique;not null;index:idx_user_username"`
	PhoneNumber      string     						`json:"phoneNumber" gorm:"unique"`
	Dob	   	  		 time.Time  						`json:"dob"`
	Password  		 string     						`json:"password" gorm:"not null"`
	Gender    		 string 	   						`json:"gender"`
	Bio				 string								`json:"bio"`	
	UserFollower	 []*UserFollower					`json:"userFollowers" gorm:"foreignKey:FollowerId;references:ID"`
	UserFollowing	 []*UserFollower					`json:"userFollowing" gorm:"foreignKey:FolloweeId;references:ID"`
	Role			 []*relation.UserRole				`json:"roles" gorm:"many2many:user_roles;"`	
	EventFavourites  []*relation.EventFavourite			`json:"eventFavourites" gorm:"many2many:event_favourites;"`
	Payments		 []*relation.UserPayment			`json:"payments" gorm:"many2many:user_payments;"`
	EventInviter     []*relation.Invitation				`json:"eventInviter" gorm:"many2many:invitations;"`
	EventInvitee     []*relation.Invitation				`json:"eventInvitee" gorm:"many2many:invitations;"`
	// DeletedAt 			gorm.DeletedAt 			`json:"deletedAt" gorm:"index"`
	// CreatedAt 			time.Time  				`json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt 			time.Time  				`json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}


func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))
	
	return nil
}

func (User) TableName() string {
	return "users"
}