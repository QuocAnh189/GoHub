package dto

type ValidateUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	FullName    string `json:"fullName" validate:"required"`
}

type ValidateUserRes struct {
	Message string `json:"Message"`
}

type SignUpReq struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	FullName    string `json:"fullName" validate:"required"`
	UserName    string `json:"userName" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type SignUpRes struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type SignInReq struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInRes struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}

type ForgotPasswordReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ForgotPasswordRes struct {
}

type ResetPasswordReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ResetPasswordRes struct {
}