package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/reviews/dto"
	"gohub/domains/reviews/model"
	"gohub/pkg/paging"
)

type IReviewRepository interface {
	ListReview(ctx context.Context, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	Create(ctx context.Context, review *model.Review) error
	Update(ctx context.Context, review *model.Review) error
	Delete(ctx context.Context, id string) error
	GetReviewByID(ctx context.Context, id string, preload bool) (*model.Review, error)
	GetReviewByEventID(ctx context.Context, eventID string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	GetReviewByUserID(ctx context.Context, userID string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error)
	GetReviewByCreatedEvents(ctx context.Context, userID string, req *dto.ListReviewReq, statistic *dto.StatisticReviewCreatedEvent) ([]*model.Review, *paging.Pagination, error)
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

func (r *ReviewRepo) GetReviewByID(ctx context.Context, id string, preload bool) (*model.Review, error) {
	var review model.Review

	opts := []database.FindOption{
		database.WithQuery(database.NewQuery("id = ?", id)),
	}
	if preload {
		opts = append(opts, database.WithPreload([]string{"User"}))
	}

	if err := r.db.FindOne(ctx, &review, opts...); err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepo) ListReview(ctx context.Context, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Content != "" {
		query = append(query, database.NewQuery("content LIKE ?", "%"+req.Content+"%"))
	}
	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &model.Review{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var reviews []*model.Review
	if err := r.db.Find(
		ctx,
		&reviews,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"User"}),
	); err != nil {
		return nil, nil, err
	}

	return reviews, pagination, nil
}

func (r *ReviewRepo) GetReviewByEventID(ctx context.Context, eventId string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Content != "" {
		query = append(query, database.NewQuery("event_id = ? AND content LIKE ?", eventId, "%"+req.Content+"%"))
	} else {
		query = append(query, database.NewQuery("event_id = ?", eventId))
	}

	order := "created_at DESC"

	var total int64
	if err := r.db.Count(ctx, &model.Review{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var reviews []*model.Review
	if err := r.db.Find(
		ctx,
		&reviews,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"User"}),
	); err != nil {
		return nil, nil, err
	}

	return reviews, pagination, nil
}

func (r *ReviewRepo) GetReviewByUserID(ctx context.Context, userID string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Content != "" {
		query = append(query, database.NewQuery("user_id = ? AND content LIKE ?", userID, "%"+req.Content+"%"))
	} else {
		query = append(query, database.NewQuery("user_id = ?", userID))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &model.Review{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var reviews []*model.Review
	if err := r.db.Find(
		ctx,
		&reviews,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return reviews, pagination, nil
}

func (r *ReviewRepo) GetReviewByCreatedEvents(ctx context.Context, userID string, req *dto.ListReviewReq, statistic *dto.StatisticReviewCreatedEvent) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Content != "" {
		query = append(query, database.NewQuery("events.user_id = ? AND content LIKE ?", userID, "%"+req.Content+"%"))
	} else {
		query = append(query, database.NewQuery("events.user_id = ?", userID))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var avgRate float64
	err := r.db.GetDB().Raw(
		`SELECT COALESCE(AVG(reviews.rate), 0) AS average_rate
		FROM reviews
		INNER JOIN events ON reviews.event_id = events.id
		WHERE events.user_id = ?`,
		userID,
	).Scan(&avgRate).Error
	if err != nil {
		return nil, nil, err
	}
	statistic.AverageRate = avgRate

	var totalPositive float64
	if err := r.db.GetDB().Raw(
		`SELECT count(*) AS total_positive
		FROM reviews
		INNER JOIN events ON reviews.event_id = events.id
		WHERE events.user_id = ? AND reviews.is_positive = ?`,
		userID, true,
	).Scan(&totalPositive).Error; err != nil {
		return nil, nil, err
	}
	statistic.TotalPositive = totalPositive

	var totalNegative float64
	if err := r.db.GetDB().Raw(
		`SELECT count(*) AS total_negative
		FROM reviews
		INNER JOIN events ON reviews.event_id = events.id
		WHERE events.user_id = ? AND reviews.is_positive = ?`,
		userID, false,
	).Scan(&totalNegative).Error; err != nil {
		return nil, nil, err
	}
	statistic.TotalNegative = totalNegative

	var rateCounts []dto.RateCount
	if err := r.db.GetDB().Raw(
		`SELECT reviews.rate, COUNT(*) AS total
		FROM reviews
		INNER JOIN events ON reviews.event_id = events.id
		WHERE events.user_id = ?
		GROUP BY reviews.rate ORDER BY reviews.rate ASC`,
		userID,
	).Scan(&rateCounts).Error; err != nil {
		return nil, nil, err
	}

	statistic.TotalPerNumberRate = rateCounts

	var total int64
	if err := r.db.Count(
		ctx,
		&model.Review{},
		&total,
		database.WithJoin("INNER JOIN events ON reviews.event_id = events.id"),
		database.WithQuery(query...),
	); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var reviews []*model.Review
	if err := r.db.Find(
		ctx,
		&reviews,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithJoin("INNER JOIN events ON reviews.event_id = events.id"),
		database.WithPreload([]string{"User", "Event"}),
	); err != nil {
		return nil, nil, err
	}

	return reviews, pagination, nil
}

func (r *ReviewRepo) Delete(ctx context.Context, id string) error {
	category, err := r.GetReviewByID(ctx, id, false)
	if err != nil {
		return err
	}
	return r.db.Delete(ctx, category)
}
