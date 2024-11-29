package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"gohub/domains/auth/dto"
	"gohub/domains/users/model"
	"gohub/domains/users/repository"
	"gohub/pkg/jwt"
	"gohub/pkg/messages"
	"gohub/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type IAuthService interface {
	ValidateUser(ctx context.Context, req *dto.ValidateUserReq) error
	SignUp(ctx context.Context, req *dto.SignUpReq) (*model.User, string, string, error)
	SignIn(ctx context.Context, req *dto.SignInReq) (*model.User, string, string, error)
	SignOut(ctx context.Context, token string) error
	ExternalSignIn(w http.ResponseWriter, r *http.Request)
	ExternalCallback(w http.ResponseWriter, r *http.Request) (goth.User, error)
	RefreshToken(ctx context.Context, userId string) (string, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, id string, req *dto.ResetPasswordReq) error
	GetProfile(ctx context.Context, id string) (*model.User, error)
}

type AuthService struct {
	validator validation.Validation
	repo      repository.IUserRepository
}

func NewAuthService(validator validation.Validation, repo repository.IUserRepository) *AuthService {
	return &AuthService{
		validator: validator,
		repo:      repo,
	}
}

func (a *AuthService) ValidateUser(ctx context.Context, req *dto.ValidateUserReq) error {
	if err := a.validator.ValidateStruct(req); err != nil {
		return err
	}

	existingEmail, err := a.repo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail != nil {
		return errors.New(messages.RegisterEmailExists)
	}

	existingUserName, err := a.repo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName != nil {
		return errors.New(messages.RegisterUserNameExists)
	}

	existingPhoneNumber, err := a.repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneNumber != nil {
		return errors.New(messages.RegisterPhoneNumberExists)
	}

	return nil
}

func (a *AuthService) SignUp(ctx context.Context, req *dto.SignUpReq) (*model.User, string, string, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return nil, "", "", err
	}

	existingEmail, err := a.repo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail != nil {
		return nil, "", "", errors.New(messages.RegisterEmailExists)
	}

	existingUserName, err := a.repo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName != nil {
		return nil, "", "", errors.New(messages.RegisterUserNameExists)
	}

	var user *model.User
	utils.MapStruct(&user, &req)
	err = a.repo.Create(ctx, user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return nil, "", "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return user, accessToken, refreshToken, nil
}

func (a *AuthService) SignIn(ctx context.Context, req *dto.SignInReq) (*model.User, string, string, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return nil, "", "", err
	}

	user, err := a.repo.GetUserByEmailOrUsername(ctx, req.Identity)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", "", errors.New(messages.AccountOrPasswordWrong)
		}
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Identity, err)
		return nil, "", "", err
	}

	fmt.Print()
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New(messages.AccountOrPasswordWrong)
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return user, accessToken, refreshToken, nil
}

func (a *AuthService) SignOut(ctx context.Context, token string) error {
	jwt.BlacklistToken(token)
	return nil
}

func (a *AuthService) ExternalSignIn(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (a *AuthService) ExternalCallback(w http.ResponseWriter, r *http.Request) (goth.User, error) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return goth.User{}, err
	}
	return user, nil
}

func (a *AuthService) RefreshToken(ctx context.Context, userId string) (string, error) {
	user, err := a.repo.GetUserByID(ctx, userId)
	if err != nil {
		logger.Errorf("RefreshToken.GetUserByID fail, id: %s, error: %s", userId, err)
		return "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}
	accessToken := jwt.GenerateAccessToken(tokenData)
	return accessToken, nil
}

func (a *AuthService) ForgotPassword(ctx context.Context, email string) error {
	panic("unimplemented")
}

func (a *AuthService) ResetPassword(ctx context.Context, id string, req *dto.ResetPasswordReq) error {
	if err := a.validator.ValidateStruct(req); err != nil {
		return err
	}
	user, err := a.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("ChangePassword.GetUserByID fail, id: %s, error: %s", id, err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New(messages.WrongPassword)
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = a.repo.Update(ctx, user)
	if err != nil {
		logger.Errorf("ChangePassword.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (a *AuthService) GetProfile(ctx context.Context, id string) (*model.User, error) {
	user, err := a.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("GetUserByID fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return user, nil
}
