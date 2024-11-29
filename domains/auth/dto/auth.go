package dto

import (
	"gohub/domains/users/model"
	"time"
)

type ValidateUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	UserName    string `json:"userName" validate:"required"`
}

type SignUpReq struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	UserName    string `json:"userName" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type SignUpRes struct {
	User         model.User `json:"user"`
	AccessToken  string     `json:"accessToken" validate:"required"`
	RefreshToken string     `json:"refreshToken" validate:"required"`
}

type SignInReq struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInRes struct {
	User         model.User `json:"user"`
	AccessToken  string     `json:"accessToken" validate:"required"`
	RefreshToken string     `json:"refreshToken" validate:"required"`
}

type SignOutRes struct {
	Message string `json:"message"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"accessToken"`
}

type ForgotPasswordReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ForgotPasswordRes struct {
}

type ResetPasswordReq struct {
	Password    string `json:"password" validate:"required,password"`
	NewPassword string `json:"new_password" validate:"required,password"`
}

type ResetPasswordRes struct {
	Message string `json:"message"`
}

type ProfileResponse struct {
	ID              string     `json:"id"`
	Email           string     `json:"email"`
	AvatarUrl       string     `json:"avatarUrl"`
	AvatarFileName  string     `json:"avatarFileName"`
	FullName        string     `json:"fullName"`
	UserName        string     `json:"userName"`
	PhoneNumber     string     `json:"phoneNumber"`
	Dob             *time.Time `json:"dob"`
	Gender          string     `json:"gender"`
	Bio             string     `json:"bio"`
	UserFollower    int32      `json:"userFollowers"`
	UserFollowing   int32      `json:"userFollowing"`
	Role            string     `json:"roles"`
	EventFavourites int32      `json:"eventFavourites"`
}
