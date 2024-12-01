package service

import (
	"context"
	"errors"
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"
	roleModel "gohub/domains/roles/model"
	roleRepository "gohub/domains/roles/repository"
	"gohub/domains/users/dto"
	"gohub/domains/users/model"
	"gohub/domains/users/repository"
	"gohub/pkg/messages"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserReq) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, *dto.Calculation, error)
	GetUsers(ctx context.Context, req *dto.ListUserReq) ([]*model.User, *paging.Pagination, error)
	UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) (*model.User, error)
	ChangePassword(ctx context.Context, id string, req *dto.ChangePassword) error
	GetFollowers(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error)
	GetFollowing(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error)
	IsFollower(ctx context.Context, req *dto.FollowerUserReq) bool
	FollowUser(ctx context.Context, req *dto.FollowerUserReq) error
	UnfollowUser(ctx context.Context, req *dto.FollowerUserReq) error
}

type UserService struct {
	validator validation.Validation
	userRepo  repository.IUserRepository
	roleRepo  roleRepository.IRoleRepository
}

func NewUserService(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	roleRepo roleRepository.IRoleRepository) *UserService {
	return &UserService{
		validator: validator,
		userRepo:  userRepo,
		roleRepo:  roleRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *dto.CreateUserReq) (*model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	existingEmail, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail != nil {
		return nil, errors.New(messages.EmailAlreadyExists)
	}

	existingUserName, err := u.userRepo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName != nil {
		return nil, errors.New(messages.UserNameAlreadyExists)
	}

	existingPhoneNumber, err := u.userRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneNumber != nil {
		return nil, errors.New(messages.PhoneNumberAlreadyExists)
	}

	var userRoles []*model.UserRole

	if len(req.Role) > 0 {
		for _, role := range req.Role {
			userRoles = append(userRoles, &model.UserRole{RoleId: role})
		}
	} else {
		var role *roleModel.Role
		role, err := u.roleRepo.GetRoleByName(ctx, "Organizer")
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, &model.UserRole{RoleId: role.ID})
	}

	var user model.User
	utils.MapStruct(&user, req)

	err = u.userRepo.CreateUser(ctx, &user, userRoles)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserService) GetUserById(ctx context.Context, id string) (*model.User, *dto.Calculation, error) {
	user, calculation, err := u.userRepo.GetUserByID(ctx, id, true)
	if err != nil {
		return nil, nil, err
	}

	return user, calculation, nil
}

func (u *UserService) GetUsers(ctx context.Context, req *dto.ListUserReq) ([]*model.User, *paging.Pagination, error) {
	users, pagination, err := u.userRepo.ListUsers(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (u *UserService) UpdateUser(ctx context.Context, id string, req *dto.UpdateUserReq) (*model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	userExists, _, err := u.userRepo.GetUserByID(ctx, id, true)
	if err != nil {
		return nil, err
	}

	existingEmail, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingEmail.Email != userExists.Email {
		return nil, errors.New(messages.EmailAlreadyExists)
	}

	existingUserName, err := u.userRepo.GetUserByUserName(ctx, req.UserName)
	if err == nil && existingUserName.UserName != userExists.UserName {
		return nil, errors.New(messages.UserNameAlreadyExists)
	}

	existingPhoneNumber, err := u.userRepo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err == nil && existingPhoneNumber.PhoneNumber != userExists.PhoneNumber {
		return nil, errors.New(messages.PhoneNumberAlreadyExists)
	}

	var user model.User
	utils.MapStruct(&user, req)
	user.Password = userExists.Password

	err = u.userRepo.UpdateUser(ctx, &user)
	if err != nil {
		logger.Errorf("Update fail, error: %s", err)
		return nil, err
	}

	user.Roles = userExists.Roles

	return &user, nil
}

func (u *UserService) ChangePassword(ctx context.Context, id string, req *dto.ChangePassword) error {
	if err := u.validator.ValidateStruct(req); err != nil {
		return err
	}
	user, _, err := u.userRepo.GetUserByID(ctx, id, false)
	if err != nil {
		logger.Errorf("ChangePassword.GetUserByID fail, id: %s, error: %s", id, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(messages.UserNotFound)
		}
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New(messages.WrongPassword)
	}

	user.Password = utils.HashAndSalt([]byte(req.NewPassword))
	err = u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		logger.Errorf("ChangePassword.Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (u *UserService) GetFollowers(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error) {
	users, pagination, err := u.userRepo.GetUserFollowers(ctx, req, id)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}

func (u *UserService) GetFollowing(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error) {
	users, pagination, err := u.userRepo.GetUserFollowings(ctx, req, id)
	if err != nil {
		return nil, nil, err
	}
	return users, pagination, nil
}

func (u *UserService) IsFollower(ctx context.Context, req *dto.FollowerUserReq) bool {
	var userFollower *model.UserFollower
	utils.MapStruct(&userFollower, &req)

	isFollower := u.userRepo.IsFollower(ctx, userFollower)
	return isFollower
}

func (u *UserService) FollowUser(ctx context.Context, req *dto.FollowerUserReq) error {
	var userFollower *model.UserFollower
	utils.MapStruct(&userFollower, &req)

	err := u.userRepo.FollowerUser(ctx, userFollower)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UnfollowUser(ctx context.Context, req *dto.FollowerUserReq) error {
	var userFollower *model.UserFollower
	utils.MapStruct(&userFollower, &req)

	err := u.userRepo.UnFollowerUser(ctx, userFollower)
	if err != nil {
		return err
	}
	return nil
}
