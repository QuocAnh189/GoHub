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
	UserName 		 string 	   						`json:"userName"`
	PhoneNumber      string     						`json:"phoneNumber"`
	Dob	   	  		 time.Time  						`json:"dob"`
	Password  		 string     						`json:"password"`
	Gender    		 int 	   							`json:"gender" gorm:"default:0"`
	Status 			 int 	   							`json:"status" gorm:"default:1"`
	UserFollower	 []*UserFollower					`json:"userFollowers" gorm:"foreignKey:FollowerId;references:ID"`
	UserFollowing	 []*UserFollower					`json:"userFollowing" gorm:"foreignKey:FollowedId;references:ID"`
	Role			 []*relation.UserRole				`json:"roles" gorm:"many2many:user_roles;"`	
	IsDeleted     	 bool       						`json:"isDeleted" gorm:"default:0"`
	EventFavourites  []*relation.EventFavourite			`json:"eventFavourites" gorm:"many2many:event_favourites;"`
	Payments		 []*relation.UserPayment			`json:"payments" gorm:"many2many:user_payments;"`
	EventInviter     []*relation.Invitation				`json:"eventInviter" gorm:"many2many:invitations;"`
	EventInvitee     []*relation.Invitation				`json:"eventInvitee" gorm:"many2many:invitations;"`
}


func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))
	
	return nil
}

func (User) TableName() string {
	return "users"
}