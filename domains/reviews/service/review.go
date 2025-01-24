package service

import (
	"context"
	"gohub/domains/reviews/dto"
	"gohub/domains/reviews/model"
	"gohub/domains/reviews/repository"
	"gohub/internal/libs/logger"
	"gohub/pkg/paging"
	"gohub/pkg/utils"

	"gohub/internal/libs/validation"
)

type IReviewService interface {
	CreateReview(ctx context.Context, req *dto.CreateReviewReq) (*model.Review, error)
	GetReviews(ctx context.Context, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	GetReviewById(ctx context.Context, id string) (*model.Review, error)
	GetReviewsByEvent(ctx context.Context, eventId string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	GetReviewsByUser(ctx context.Context, userId string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	GetReviewByCreatedEvents(ctx context.Context, userId string, req *dto.ListReviewReq, statistic *dto.StatisticReviewCreatedEvent) ([]*model.Review, *paging.Pagination, error)
	UpdateReview(ctx context.Context, id string, req *dto.UpdateReviewReq) (*model.Review, error)
	DeleteReview(ctx context.Context, id string) error
}

type ReviewService struct {
	validator  validation.Validation
	repoReview repository.IReviewRepository
}

func NewReviewService(validator validation.Validation, repo repository.IReviewRepository) *ReviewService {
	return &ReviewService{
		validator:  validator,
		repoReview: repo,
	}
}

func (r *ReviewService) CreateReview(ctx context.Context, req *dto.CreateReviewReq) (*model.Review, error) {
	if err := r.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var review model.Review
	utils.MapStruct(&review, req)
	logger.Info(review.IsPositive)

	err := r.repoReview.Create(ctx, &review)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &review, nil
}

func (r *ReviewService) GetReviews(ctx context.Context, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	reviews, pagination, err := r.repoReview.ListReview(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (r *ReviewService) GetReviewById(ctx context.Context, id string) (*model.Review, error) {
	review, err := r.repoReview.GetReviewByID(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewService) GetReviewsByEvent(ctx context.Context, eventId string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	reviews, pagination, err := r.repoReview.GetReviewByEventID(ctx, eventId, req)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (r *ReviewService) GetReviewsByUser(ctx context.Context, userId string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	reviews, pagination, err := r.repoReview.GetReviewByUserID(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (r *ReviewService) GetReviewByCreatedEvents(ctx context.Context, userId string, req *dto.ListReviewReq, statistic *dto.StatisticReviewCreatedEvent) ([]*model.Review, *paging.Pagination, error) {
	reviews, pagination, err := r.repoReview.GetReviewByCreatedEvents(ctx, userId, req, statistic)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (r *ReviewService) UpdateReview(ctx context.Context, id string, req *dto.UpdateReviewReq) (*model.Review, error) {
	if err := r.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	review, err := r.repoReview.GetReviewByID(ctx, id, false)
	if err != nil {
		logger.Errorf("Update.GetCategoryByID fail, id: %s, error: %s", id, err)
	}

	utils.MapStruct(review, req)
	err = r.repoReview.Update(ctx, review)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return review, nil
}

func (r *ReviewService) DeleteReview(ctx context.Context, id string) error {
	err := r.repoReview.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
