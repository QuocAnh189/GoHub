package dto

import (
	"gohub/pkg/paging"
)

type InfoFollow struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	AvatarUrl      string `json:"avatarUrl"`
	AvatarFileName string `json:"avatarFileName"`
	FullName       string `json:"fullName"`
	UserName       string `json:"userName"`
}

type User struct {
	ID             string  `json:"id"`
	Email          string  `json:"email"`
	AvatarUrl      string  `json:"avatarUrl"`
	AvatarFileName string  `json:"avatarFileName"`
	FullName       string  `json:"fullName"`
	UserName       string  `json:"userName"`
	PhoneNumber    string  `json:"phoneNumber"`
	Dob            string  `json:"dob"`
	Gender         string  `json:"gender"`
	Bio            string  `json:"bio"`
	TotalFollower  int64   `json:"totalFollower"`
	TotalFollowing int64   `json:"totalFollowing"`
	Roles          []*Role `json:"roles"`
}

type CreateUserReq struct {
	Email          string   `form:"email" validate:"required,email"`
	AvatarUrl      string   `form:"avatarUrl"`
	AvatarFileName string   `form:"avatarFileName"`
	FullName       string   `form:"fullName"`
	UserName       string   `form:"userName" validate:"required"`
	PhoneNumber    string   `form:"phoneNumber" validate:"required"`
	Dob            string   `form:"dob"`
	Gender         string   `form:"gender"`
	Bio            string   `form:"bio"`
	Password       string   `form:"password"`
	Role           []string `form:"roles"`
}

type UpdateUserReq struct {
	ID             string   `form:"id"`
	Email          string   `form:"email" validate:"required,email"`
	AvatarUrl      string   `form:"avatarUrl"`
	AvatarFileName string   `form:"avatarFileName"`
	FullName       string   `form:"fullName"`
	UserName       string   `form:"userName" validate:"required"`
	PhoneNumber    string   `form:"phoneNumber" validate:"required"`
	Dob            string   `form:"dob"`
	Gender         string   `form:"gender"`
	Bio            string   `form:"bio"`
	Role           []string `form:"roles"`
}

type UserList struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	AvatarUrl string  `json:"avatarUrl"`
	UserName  string  `json:"userName"`
	Roles     []*Role `json:"roles"`
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
	NewPassword string `json:"new_password" validate:"required,password"`
}
