package service

import (
	"context"
	"gohub/domains/reviews/model"
	"gohub/domains/reviews/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IReviewService interface {
	CreateReview(ctx context.Context, review *model.Review) (*model.Review, error)
	GetReviews(ctx context.Context, ids []string) ([]*model.Review, error)
	GetReview(ctx context.Context, id string) (*model.Review, error)
	GetReviewsByEvent(ctx context.Context, id string) ([]*model.Review, error)
	GetReviewsByUser(ctx context.Context, id string) ([]*model.Review, error)
	UpdateReview(ctx context.Context, review *model.Review) (*model.Review, error)
	DeleteReview(ctx context.Context, id string) error
}


type ReviewService struct {
	validator validation.Validation
	repo      repository.IReviewRepository
}

func NewReviewService(validator validation.Validation, repo repository.IReviewRepository) *ReviewService {
	return &ReviewService{
		validator: validator,
		repo:      repo,
	}
}

func (r *ReviewService) CreateReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	panic("unimplemented")
}

func (r *ReviewService) GetReviews(ctx context.Context, ids []string) ([]*model.Review, error) {
	panic("unimplemented")
}

func (r *ReviewService) GetReview(ctx context.Context, id string) (*model.Review, error) {
    panic("unimplemented")
}

func (r *ReviewService) GetReviewsByEvent(ctx context.Context, id string) ([]*model.Review, error) {
	panic("unimplemented")
}

func (r *ReviewService) GetReviewsByUser(ctx context.Context, id string) ([]*model.Review, error) {
    panic("unimplemented")
}

func (r *ReviewService) UpdateReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	panic("unimplemented")
}

func (r *ReviewService) DeleteReview(ctx context.Context, id string) error {
	panic("unimplemented")
}

