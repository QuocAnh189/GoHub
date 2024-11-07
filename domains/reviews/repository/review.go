package repository

import (
	"context"
	"gohub/database"
	"gohub/domains/reviews/model"
)

type IReviewRepository interface {
	Create(ctx context.Context, review *model.Review) error
	Update(ctx context.Context, review *model.Review) error
	GetReviewByID(ctx context.Context, id string) (*model.Review, error)
	GetReviewByUserID(ctx context.Context, userID string) ([]model.Review, error)
	GetReviewByProductID(ctx context.Context, productID string) ([]model.Review, error)
	GetReviewByRating(ctx context.Context, rating int) ([]model.Review, error)
	GetReviewByComment(ctx context.Context, comment string) ([]model.Review, error)
}

type ReviewRepo struct {
	db database.IDatabase
}

func NewReviewRepository(db database.IDatabase) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) Create(ctx context.Context, review *model.Review) error {
	return r.db.Create(ctx, review)
}

func (r *ReviewRepo) Update(ctx context.Context, review *model.Review) error {
	return r.db.Update(ctx, review)
}

func (r *ReviewRepo) GetReviewByID(ctx context.Context, id string) (*model.Review, error) {
	var review model.Review
	if err := r.db.FindById(ctx, id, &review); err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepo) GetReviewByUserID(ctx context.Context, userID string) ([]model.Review, error) {
	var reviews []model.Review
	query := database.NewQuery("user_id = ?", userID)
	if err := r.db.Find(ctx, &reviews, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepo) GetReviewByProductID(ctx context.Context, productID string) ([]model.Review, error) {
	var reviews []model.Review
	query := database.NewQuery("product_id = ?", productID)
	if err := r.db.Find(ctx, &reviews, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepo) GetReviewByRating(ctx context.Context, rating int) ([]model.Review, error) {
	var reviews []model.Review
	query := database.NewQuery("rating = ?", rating)
	if err := r.db.Find(ctx, &reviews, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepo) GetReviewByComment(ctx context.Context, comment string) ([]model.Review, error) {	
	var reviews []model.Review
	query := database.NewQuery("comment = ?", comment)
	if err := r.db.Find(ctx, &reviews, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return reviews, nil
}