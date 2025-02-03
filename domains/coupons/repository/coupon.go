package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/coupons/dto"
	"gohub/domains/coupons/model"
	"gohub/pkg/paging"
)

type ICouponRepository interface {
	ListCoupons(ctx context.Context, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error)
	GetCreatedCoupons(ctx context.Context, userId string, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error)
	GetCouponById(ctx context.Context, id string) (*model.Coupon, error)
	GetCouponByNameAndUserId(ctx context.Context, userId string, name string) (*model.Coupon, error)
	Create(ctx context.Context, coupon *model.Coupon) error
	Update(ctx context.Context, coupon *model.Coupon) error
	Delete(ctx context.Context, id string) error
}

type CouponRepository struct {
	db database.IDatabase
}

func NewCouponRepository(db database.IDatabase) *CouponRepository {
	return &CouponRepository{db: db}
}

func (c *CouponRepository) ListCoupons(ctx context.Context, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Search != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Search+"%"))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := c.db.Count(ctx, &model.Coupon{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var coupons []*model.Coupon
	if err := c.db.Find(
		ctx,
		&coupons,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return coupons, pagination, nil
}

func (c *CouponRepository) GetCreatedCoupons(ctx context.Context, userId string, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Search != "" {
		query = append(query, database.NewQuery("user_id = ? AND name LIKE ?", userId, "%"+req.Search+"%"))
	} else {
		query = append(query, database.NewQuery("user_id = ? ", userId))
	}

	order := "created_at DESC"

	var total int64
	if err := c.db.Count(ctx, &model.Coupon{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var coupons []*model.Coupon
	if err := c.db.Find(
		ctx,
		&coupons,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return coupons, pagination, nil
}

func (c *CouponRepository) GetCouponById(ctx context.Context, id string) (*model.Coupon, error) {
	var coupon model.Coupon
	if err := c.db.FindById(ctx, id, &coupon); err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (c *CouponRepository) GetCouponByNameAndUserId(ctx context.Context, userId string, name string) (*model.Coupon, error) {
	var coupon model.Coupon
	query := database.NewQuery("name = ? AND user_id = ?", name, userId)
	if err := c.db.FindOne(ctx, &coupon, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (c *CouponRepository) Create(ctx context.Context, coupon *model.Coupon) error {
	return c.db.Create(ctx, coupon)
}

func (c *CouponRepository) Update(ctx context.Context, coupon *model.Coupon) error {
	return c.db.Update(ctx, coupon)
}

func (c *CouponRepository) Delete(ctx context.Context, id string) error {
	coupon, err := c.GetCouponById(ctx, id)
	if err != nil {
		return err
	}
	return c.db.Delete(ctx, coupon)
}
