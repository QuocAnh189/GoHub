package dto

import (
	"gohub/pkg/paging"
	"mime/multipart"
)

type InfoFollow struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	FullName  string `json:"fullName"`
	UserName  string `json:"userName"`
}

type User struct {
	ID             string  `json:"id"`
	Email          string  `json:"email"`
	AvatarUrl      string  `json:"avatarUrl"`
	FullName       string  `json:"fullName"`
	UserName       string  `json:"userName"`
	PhoneNumber    string  `json:"phoneNumber"`
	Dob            string  `json:"dob"`
	Gender         string  `json:"gender"`
	Bio            string  `json:"bio"`
	TotalEvent     int64   `json:"totalEvent"`
	TotalFollower  int64   `json:"totalFollower"`
	TotalFollowing int64   `json:"totalFollowing"`
	Roles          []*Role `json:"roles"`
}

type CreateUserReq struct {
	Email       string   `form:"email" validate:"required,email"`
	AvatarUrl   string   `form:"avatarUrl"`
	FullName    string   `form:"fullName"`
	UserName    string   `form:"userName" validate:"required"`
	PhoneNumber string   `form:"phoneNumber" validate:"required"`
	Dob         string   `form:"dob"`
	Gender      string   `form:"gender"`
	Bio         string   `form:"bio"`
	Password    string   `form:"password"`
	Role        []string `form:"roles"`
}

type UpdateUserReq struct {
	ID          string                `form:"id"`
	Email       string                `form:"email" validate:"required,email"`
	AvatarUrl   string                `form:"avatarUrl"`
	Avatar      *multipart.FileHeader `form:"avatar"`
	FullName    string                `form:"fullName"`
	UserName    string                `form:"userName" validate:"required"`
	PhoneNumber string                `form:"phoneNumber" validate:"required"`
	Dob         string                `form:"dob"`
	Gender      string                `form:"gender"`
	Bio         string                `form:"bio"`
}

type UserList struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	AvatarUrl    string  `json:"avatarUrl"`
	UserName     string  `json:"userName"`
	IsInvitation bool    `json:"isInvitation"`
	Roles        []*Role `json:"roles"`
}

type Invitations struct {
	ID        string          `json:"id"`
	Inviter   Inviter         `json:"inviter"`
	InviteeId string          `json:"inviteeId"`
	Event     EventInvitation `json:"event"`
	CreatedAt string          `json:"createdAt"`
}

type Inviter struct {
	ID        string `json:"id"`
	AvatarUrl string `json:"avatarUrl"`
	FullName  string `json:"fullName"`
}

type EventInvitation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NotificationFollowing struct {
	ID         string     `json:"id"`
	Follower   InfoFollow `json:"follower"`
	FolloweeId string     `json:"followeeId"`
	CreatedAt  string     `json:"createdAt"`
}

type Calculation struct {
	TotalEvent     int64 `json:"totalEvent"`
	TotalFollower  int64 `json:"totalFollower"`
	TotalFollowing int64 `json:"totalFollowing"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListUserReq struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListUserRes struct {
	User       []*UserList        `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type FollowerUserReq struct {
	FollowerId string `json:"followerId"`
	FolloweeId string `json:"followeeId"`
}

type ChangePassword struct {
	Password    string `json:"password" validate:"required,password"`
	NewPassword string `json:"newPassword" validate:"required,password"`
}

type InviteUsers struct {
	UserIds []string `json:"userIds"`
	EventId string   `json:"eventId"`
}

type ListInvitationReq struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListInvitationRes struct {
	Invitations []*Invitations     `json:"items"`
	Pagination  *paging.Pagination `json:"metadata"`
}

type ListNotificationReq struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListNotificationFollowingRes struct {
	Notifications []*NotificationFollowing `json:"items"`
	Pagination    *paging.Pagination       `json:"metadata"`
}

type CheckInvitationReq struct {
	InviteeId string `json:"-" form:"inviteeId"`
	EventId   string `json:"-" form:"eventId"`
}
