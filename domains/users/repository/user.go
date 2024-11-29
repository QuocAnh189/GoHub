package repository

import (
	"context"
	"gohub/database"
	"gohub/domains/users/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmailOrUsername(ctx context.Context, identity string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error)
}

type UserRepo struct {
	db database.IDatabase
}

func NewUserRepository(db database.IDatabase) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	return r.db.Create(ctx, user)
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	return r.db.Update(ctx, user)
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.FindById(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByEmailOrUsername(ctx context.Context, identity string) (*model.User, error) {
	var user model.User
	queryEmail := database.NewQuery("email = ?", identity)
	queryUsername := database.NewQuery("user_name = ?", identity)
	if err := r.db.FindOne(ctx, &user, database.WithQuery(queryEmail, queryUsername)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("email = ?", email)
	if err := r.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("user_name = ?", username)
	if err := r.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("phone_number = ?", phoneNumber)
	if err := r.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}
