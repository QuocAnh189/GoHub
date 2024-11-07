package service

import (
	"context"
	"gohub/domains/users/model"
	"gohub/domains/users/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *model.User)
	GetUser(ctx context.Context, id string)
	GetUsers(ctx context.Context, ids []string)
	UpdateUser(ctx context.Context, user *model.User)
	ChangePassword(ctx context.Context)
	GetFollowers(ctx context.Context)
	GetFollowing(ctx context.Context)
	FollowUser(ctx context.Context)
	UnfollowUser(ctx context.Context)
}


type UserService struct {
	validator validation.Validation
	repo 	repository.IUserRepository
}

func NewUserService(validator validation.Validation, repo repository.IUserRepository) *UserService {
	return &UserService{
		validator: validator,
		repo: repo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *model.User) {
	panic("unimplemented")
}

func (u *UserService) GetUser(ctx context.Context, id string) {
	panic("unimplemented")
}

func (u *UserService) GetUsers(ctx context.Context, ids []string) {
    panic("unimplemented")
}

func (u *UserService) UpdateUser(ctx context.Context, user *model.User) {
    panic("unimplemented")
}

func (u *UserService) ChangePassword(ctx context.Context) {
    panic("unimplemented")
}

func (u *UserService) GetFollowers(ctx context.Context) {
    panic("unimplemented")
}

func (u *UserService) GetFollowing(ctx context.Context) {
    panic("unimplemented")
}

func (u *UserService) FollowUser(ctx context.Context) {
	panic("unimplemented")
}

func (u *UserService) UnfollowUser(ctx context.Context) {
	
}