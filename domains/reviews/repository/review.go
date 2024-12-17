package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/reviews/dto"
	"gohub/domains/reviews/model"
	"gohub/pkg/paging"
	"math"
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
	if req.Search != "" {
		query = append(query, database.NewQuery("content LIKE ?", "%"+req.Search+"%"))
	}

	order := "created_at DESC"
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
	if req.Search != "" {
		query = append(query, database.NewQuery("event_id = ? AND content LIKE ?", eventId, "%"+req.Search+"%"))
	} else {
		query = append(query, database.NewQuery("event_id = ?", eventId))
	}

	order := "created_at DESC"
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

func (r *ReviewRepo) GetReviewByUserID(ctx context.Context, userID string, req *dto.ListReviewReq) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Search != "" {
		query = append(query, database.NewQuery("user_id = ? AND content LIKE ?", userID, "%"+req.Search+"%"))
	} else {
		query = append(query, database.NewQuery("user_id = ?", userID))
	}

	order := "created_at DESC"
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

func (r *ReviewRepo) GetReviewByCreatedEvents(ctx context.Context, userId string, req *dto.ListReviewReq, statistic *dto.StatisticReviewCreatedEvent) ([]*model.Review, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "events.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND content ILIKE ? OR users.user_name ILIKE ? OR users.full_name ILIKE ? OR users.email ILIKE ?"
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
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
		userId,
	).Scan(&avgRate).Error
	if err != nil {
		return nil, nil, err
	}
	statistic.AverageRate = math.Round(avgRate*100) / 100

	var totalPositive float64
	if err := r.db.GetDB().Raw(
		`SELECT count(*) AS total_positive
		FROM reviews
		INNER JOIN events ON reviews.event_id = events.id
		WHERE events.user_id = ? AND reviews.is_positive = ?`,
		userId, true,
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
		userId, false,
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
		userId,
	).Scan(&rateCounts).Error; err != nil {
		return nil, nil, err
	}

	statistic.TotalPerNumberRate = rateCounts

	var total int64
	if err := r.db.Count(
		ctx,
		&model.Review{},
		&total,
		database.WithJoin(`
			INNER JOIN events ON reviews.event_id = events.id
			INNER JOIN users ON users.id = reviews.user_id
		`),
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
		database.WithJoin(`
			INNER JOIN events ON reviews.event_id = events.id
			INNER JOIN users ON users.id = reviews.user_id
		`),
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
