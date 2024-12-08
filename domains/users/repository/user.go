package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	modelEvent "gohub/domains/events/model"
	roleModel "gohub/domains/roles/model"
	"gohub/domains/users/dto"
	"gohub/domains/users/model"
	"gohub/pkg/paging"
	"strings"
)

type IUserRepository interface {
	ListUsers(ctx context.Context, req *dto.ListUserReq) ([]*model.User, *paging.Pagination, error)
	CreateUser(ctx context.Context, user *model.User, userRoles []*model.UserRole) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string, preload bool) (*model.User, *dto.Calculation, error)
	GetUserByEmailOrUsername(ctx context.Context, identity string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error)
	GetUserFollowers(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error)
	GetUserFollowings(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error)
	IsFollower(ctx context.Context, userFollower *model.UserFollower) bool
	FollowerUser(ctx context.Context, userFollower *model.UserFollower) error
	UnFollowerUser(ctx context.Context, userFollower *model.UserFollower) error
}

type UserRepository struct {
	db database.IDatabase
}

func NewUserRepository(db database.IDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) ListUsers(ctx context.Context, req *dto.ListUserReq) ([]*model.User, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Name != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}
	if req.Email != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := u.db.Count(ctx, &model.User{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var users []*model.User
	if err := u.db.Find(
		ctx,
		&users,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Roles"}),
	); err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.User, userRoles []*model.UserRole) error {
	handler := func() error {
		if err := u.db.Create(ctx, user); err != nil {
			return err
		}

		var roleIds []string
		for _, userRole := range userRoles {
			userRole.UserId = user.ID
			roleIds = append(roleIds, userRole.RoleId)
		}

		if err := u.db.CreateInBatches(ctx, &userRoles, len(userRoles)); err != nil {
			return err
		}

		var roles []*roleModel.Role
		query := database.NewQuery("id IN ?", roleIds)

		err := u.db.Find(ctx, &roles, database.WithQuery(query))
		if err != nil {
			return err
		}

		user.Roles = roles

		return nil
	}

	err := u.db.WithTransaction(handler)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return u.db.Update(ctx, user)
}

func (u *UserRepository) GetUserByID(ctx context.Context, id string, preload bool) (*model.User, *dto.Calculation, error) {
	var user model.User
	var calculation dto.Calculation
	opts := []database.FindOption{
		database.WithQuery(database.NewQuery("id = ?", id)),
	}

	if preload {
		opts = append(opts, database.WithPreload([]string{"Roles"}))
	}

	var totalEvent int64
	if err := u.db.Count(ctx, &modelEvent.Event{}, &totalEvent, database.WithQuery(database.NewQuery("user_id = ?", id))); err != nil {
		return nil, nil, err
	}

	var totalFollower int64
	if err := u.db.Count(ctx, &model.UserFollower{}, &totalFollower, database.WithQuery(database.NewQuery("followee_id = ?", id))); err != nil {
		return nil, nil, err
	}

	var totalFollowing int64
	if err := u.db.Count(ctx, &model.UserFollower{}, &totalFollowing, database.WithQuery(database.NewQuery("follower_id = ?", id))); err != nil {
		return nil, nil, err
	}

	if err := u.db.FindOne(ctx, &user, opts...); err != nil {
		return nil, nil, err
	}

	calculation.TotalEvent = totalEvent
	calculation.TotalFollower = totalFollower
	calculation.TotalFollowing = totalFollowing

	return &user, &calculation, nil
}

func (u *UserRepository) GetUserByEmailOrUsername(ctx context.Context, identity string) (*model.User, error) {
	var user model.User
	queryEmail := database.NewQuery("email = ?", identity)
	queryUsername := database.NewQuery("user_name = ?", identity)
	if err := u.db.FindOne(ctx, &user, database.WithQuery(queryEmail, queryUsername)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("email = ?", email)
	if err := u.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("user_name = ?", username)
	if err := u.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	var user model.User
	query := database.NewQuery("phone_number = ?", phoneNumber)
	if err := u.db.FindOne(ctx, &user, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserFollowers(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	query = append(query, database.NewQuery("followee_id = ?", id))

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := u.db.Count(ctx, &model.UserFollower{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var users []*model.UserFollower
	if err := u.db.Find(
		ctx,
		&users,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Follower", "Follower.Roles"}),
	); err != nil {
		return nil, nil, err
	}

	var results []*model.User
	for _, user := range users {
		results = append(results, user.Follower)
	}

	if req.Name != "" || req.Email != "" {
		var results []*model.User
		for _, user := range users {
			if (req.Name != "" && strings.Contains(user.Follower.UserName, req.Name)) ||
				(req.Email != "" && strings.Contains(user.Follower.Email, req.Email)) {
				results = append(results, user.Follower)
			}
		}

		pagination := paging.NewPagination(req.Page, req.Limit, int64(len(results)))

		return results, pagination, nil
	}

	return results, pagination, nil
}

func (u *UserRepository) GetUserFollowings(ctx context.Context, req *dto.ListUserReq, id string) ([]*model.User, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	query = append(query, database.NewQuery("follower_id = ?", id))

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := u.db.Count(ctx, &model.UserFollower{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var users []*model.UserFollower
	if err := u.db.Find(
		ctx,
		&users,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"Followee", "Followee.Roles"}),
	); err != nil {
		return nil, nil, err
	}

	var results []*model.User
	for _, user := range users {
		results = append(results, user.Followee)
	}

	if req.Name != "" || req.Email != "" {
		var results []*model.User
		for _, user := range users {
			if (req.Name != "" && strings.Contains(user.Followee.UserName, req.Name)) ||
				(req.Email != "" && strings.Contains(user.Followee.Email, req.Email)) {
				results = append(results, user.Followee)
			}
		}

		pagination := paging.NewPagination(req.Page, req.Limit, int64(len(results)))

		return results, pagination, nil
	}

	return results, pagination, nil
}

func (u *UserRepository) IsFollower(ctx context.Context, userFollower *model.UserFollower) bool {
	query := database.NewQuery("follower_id = ? AND followee_id = ?", userFollower.FollowerId, userFollower.FolloweeId)
	if err := u.db.FindOne(ctx, &userFollower, database.WithQuery(query)); err != nil {
		return false
	}
	return true
}

func (u *UserRepository) FollowerUser(ctx context.Context, userFollower *model.UserFollower) error {
	return u.db.Create(ctx, userFollower)
}

func (u *UserRepository) UnFollowerUser(ctx context.Context, userFollower *model.UserFollower) error {
	query := database.NewQuery("follower_id = ? AND followee_id = ?", userFollower.FollowerId, userFollower.FolloweeId)

	if err := u.db.ForceDelete(ctx, &userFollower, database.WithQuery(query)); err != nil {
		return err
	}
	return nil
}
