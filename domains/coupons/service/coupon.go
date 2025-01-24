package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gohub/domains/coupons/dto"
	"gohub/domains/coupons/model"
	"gohub/domains/coupons/repository"
	"gohub/internal/libs/logger"
	"gohub/internal/libs/validation"
	"gohub/pkg/messages"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
)

type ICouponService interface {
	CreateCoupon(ctx context.Context, req *dto.CreateCouponReq) (*model.Coupon, error)
	GetCoupons(ctx context.Context, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error)
	GetCreatedCoupons(ctx context.Context, userId string, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error)
	GetCouponById(ctx context.Context, id string) (*model.Coupon, error)
	DeleteCoupon(ctx context.Context, id string) error
	UpdateCoupon(ctx context.Context, id string, req *dto.UpdateCouponReq) (*model.Coupon, error)
}

type CouponService struct {
	validator  validation.Validation
	repoCoupon repository.ICouponRepository
}

func NewCouponService(validator validation.Validation, repoCoupon repository.ICouponRepository) *CouponService {
	return &CouponService{
		validator:  validator,
		repoCoupon: repoCoupon,
	}
}

func (s *CouponService) CreateCoupon(ctx context.Context, req *dto.CreateCouponReq) (*model.Coupon, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var coupon model.Coupon
	utils.MapStruct(&coupon, req)

	if req.Image.Header != nil && req.Image.Filename != "" {
		uploadUrl, err := utils.ImageUpload(req.Image, "/eventhub/conpons")
		if err != nil {
			return nil, err
		}

		coupon.CoverImageFileName = req.Image.Filename
		coupon.CoverImageUrl = uploadUrl
	}

	err := s.repoCoupon.Create(ctx, &coupon)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, errors.New(messages.CouponNameAlreadyExists)
		}
		return nil, errors.New("some thing went wrong")
	}

	return &coupon, nil
}

func (s *CouponService) GetCoupons(ctx context.Context, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error) {
	coupons, pagination, err := s.repoCoupon.ListCoupons(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return coupons, pagination, nil
}

func (s *CouponService) GetCreatedCoupons(ctx context.Context, userId string, req *dto.ListCouponReq) ([]*model.Coupon, *paging.Pagination, error) {
	coupons, pagination, err := s.repoCoupon.GetCreatedCoupons(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return coupons, pagination, nil
}

func (s *CouponService) GetCouponById(ctx context.Context, id string) (*model.Coupon, error) {
	coupon, err := s.repoCoupon.GetCouponById(ctx, id)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (s *CouponService) DeleteCoupon(ctx context.Context, id string) error {
	err := s.repoCoupon.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *CouponService) UpdateCoupon(ctx context.Context, id string, req *dto.UpdateCouponReq) (*model.Coupon, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	coupon, err := s.repoCoupon.GetCouponById(ctx, id)
	if err != nil {
		logger.Errorf("Update.GetCategoryByID fail, id: %s, error: %s", id, err)
		return nil, errors.New(messages.CategoryNotFound)
	}

	utils.MapStruct(coupon, req)
	if req.Image.Header != nil && req.Image.Filename != "" {
		logger.Info("vao day")
		uploadUrl, err := utils.ImageUpload(req.Image, "/eventhub/conpons")
		if err != nil {
			return nil, err
		}

		coupon.CoverImageFileName = req.Image.Filename
		coupon.CoverImageUrl = uploadUrl
	}

	err = s.repoCoupon.Update(ctx, coupon)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, errors.New(messages.CouponNameAlreadyExists)
		}
		return nil, errors.New("some thing went wrong")
	}

	return coupon, nil
}
