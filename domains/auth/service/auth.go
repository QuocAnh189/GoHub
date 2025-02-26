package service

import (
	"context"
	"errors"
	"gohub/domains/auth/dto"
	roleModel "gohub/domains/roles/model"
	roleRepository "gohub/domains/roles/repository"
	"gohub/domains/users/model"
	"gohub/domains/users/repository"
	"gohub/internal/libs/logger"
	"gohub/internal/libs/validation"
	"gohub/pkg/jwt"
	"gohub/pkg/messages"
	"gohub/pkg/utils"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	ValidateUser(ctx context.Context, req *dto.ValidateUserReq) error
	SignUp(ctx context.Context, req *dto.SignUpReq) (string, string, error)
	SignIn(ctx context.Context, req *dto.SignInReq) (string, string, error)
	SignOut(ctx context.Context, token string) error
	ExternalSignIn(w http.ResponseWriter, r *http.Request)
	ExternalCallback(w http.ResponseWriter, r *http.Request) (goth.User, error)
	RefreshToken(ctx context.Context, userId string) (string, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, id string, req *dto.ResetPasswordReq) error
}

type AuthService struct {
	validator validation.Validation
	authRepo  repository.IUserRepository
	roleRepo  roleRepository.IRoleRepository
}

func NewAuthService(
	validator validation.Validation,
	authRepo repository.IUserRepository,
	roleRepo roleRepository.IRoleRepository) *AuthService {
	return &AuthService{
		validator: validator,
		authRepo:  authRepo,
		roleRepo:  roleRepo,
	}
}

func (a *AuthService) ValidateUser(ctx context.Context, req *dto.ValidateUserReq) error {
	if err := a.validator.ValidateStruct(req); err != nil {
		return err
	}

	existingEmail, err := a.authRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail != nil {
		return errors.New(messages.EmailAlreadyExists)
	}

	existingUserName, err := a.authRepo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName != nil {
		return errors.New(messages.UserNameAlreadyExists)
	}

	existingPhoneNumber, err := a.authRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneNumber != nil {
		return errors.New(messages.PhoneNumberAlreadyExists)
	}

	return nil
}

func (a *AuthService) SignUp(ctx context.Context, req *dto.SignUpReq) (string, string, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return "", "", err
	}

	existingEmail, err := a.authRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail != nil {
		return "", "", errors.New(messages.EmailAlreadyExists)
	}

	existingUserName, err := a.authRepo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName != nil {
		return "", "", errors.New(messages.UserNameAlreadyExists)
	}

	existingPhoneNumber, err := a.authRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneNumber != nil {
		return "", "", errors.New(messages.PhoneNumberAlreadyExists)
	}

	var userRoles []*model.UserRole
	var role *roleModel.Role
	role, err = a.roleRepo.GetRoleByName(ctx, "Organizer")
	if err != nil {
		return "", "", err
	}
	userRoles = append(userRoles, &model.UserRole{RoleId: role.ID})

	var user *model.User
	utils.MapStruct(&user, &req)
	err = a.authRepo.CreateUser(ctx, user, userRoles)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return "", "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, nil
}

func (a *AuthService) SignIn(ctx context.Context, req *dto.SignInReq) (string, string, error) {
	if err := a.validator.ValidateStruct(req); err != nil {
		return "", "", err
	}

	user, err := a.authRepo.GetUserByEmailOrUsername(ctx, req.Identity)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", errors.New(messages.AccountOrPasswordWrong)
		}
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Identity, err)
		return "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", errors.New(messages.AccountOrPasswordWrong)
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, nil
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
	user, _, err := a.authRepo.GetUserByID(ctx, userId, false)
	if err != nil {
		logger.Errorf("RefreshToken.GetUserByID fail, id: %s, error: %s", userId, err)
		return "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
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
	user, _, err := a.authRepo.GetUserByID(ctx, id, false)
	if err != nil {
		logger.Errorf("ChangePassword.GetUserByID fail, id: %s, error: %s", id, err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New(messages.WrongPassword)
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = a.authRepo.UpdateUser(ctx, user)
	if err != nil {
		logger.Errorf("ChangePassword.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}
