package dto

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
