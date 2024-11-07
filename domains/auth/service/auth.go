package service

import (
	"context"
	"gohub/domains/users/model"
	"gohub/domains/users/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IAuthService interface {
	SignUp(ctx context.Context, user *model.User) error
	ValidateUser(ctx context.Context, user *model.User) error
	SignIn(ctx context.Context, email, password string) (string, error)
	SignOut(ctx context.Context, token string) error
	ExternalSignIn(ctx context.Context, email string) (string, error)
	ExternalCallback(ctx context.Context, email, token string) (string, error)
	RefreshToken(ctx context.Context, token string) (string, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, password string) error
	GetProfile(ctx context.Context, token string) (*model.User, error)
}

type AuthService struct {
	validator validation.Validation
	repo      repository.IUserRepository
}

func NewAuthService(
	validator validation.Validation,
	repo repository.IUserRepository) *AuthService {
	return &AuthService{
		validator: validator,
		repo:      repo,
	}
}

// ValidateUser implements IAuthService.
func (a *AuthService) ValidateUser(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

// SignUp implements IAuthService.
func (a *AuthService) SignUp(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

// SignIn implements IAuthService.
func (a *AuthService) SignIn(ctx context.Context, email string, password string) (string, error) {
	panic("unimplemented")
}

// SignOut implements IAuthService.
func (a *AuthService) SignOut(ctx context.Context, token string) error {
	panic("unimplemented")
}

// ExternalSignIn implements IAuthService.
func (a *AuthService) ExternalSignIn(ctx context.Context, email string) (string, error) {
	panic("unimplemented")
}

// ExternalCallback implements IAuthService.
func (a *AuthService) ExternalCallback(ctx context.Context, email string, token string) (string, error) {
	panic("unimplemented")
}

// RefreshToken implements IAuthService.
func (a *AuthService) RefreshToken(ctx context.Context, token string) (string, error) {
	panic("unimplemented")
}

// ForgotPassword implements IAuthService.
func (a *AuthService) ForgotPassword(ctx context.Context, email string) error {
	panic("unimplemented")
}

// ResetPassword implements IAuthService.
func (a *AuthService) ResetPassword(ctx context.Context, token string, password string) error {
	panic("unimplemented")
}

// GetProfile implements IAuthService.
func (a *AuthService) GetProfile(ctx context.Context, token string) (*model.User, error) {
	panic("unimplemented")
}
