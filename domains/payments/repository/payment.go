package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/payments/dto"
	"gohub/domains/payments/model"
	"gohub/pkg/paging"
)

type IPaymentRepository interface {
	GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error)
	GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error)
}

type PaymentRepository struct {
	db database.IDatabase
}

func NewPaymentRepository(db database.IDatabase) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (p *PaymentRepository) GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "events.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND (customer_name ILIKE ? OR events.name ILIKE ?)"
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.StartDate != "" && req.EndDate != "" {
		startDate := req.StartDate + " 00:00:00"
		endDate := req.EndDate + " 23:59:59"
		queryString += " AND payments.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := p.db.Count(
		ctx,
		&model.Payment{},
		&total,
		database.WithJoin(`
			INNER JOIN events ON payments.event_id = events.id
    	`),
		database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var transactions []*model.Payment
	if err := p.db.Find(
		ctx,
		&transactions,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithJoin(`
			INNER JOIN events ON payments.event_id = events.id
    	`),
		database.WithPreload([]string{"Event"}),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return transactions, pagination, nil
}

func (p *PaymentRepository) GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "payments.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND (events.name ILIKE ? OR users.user_name ILIKE ?)"
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if req.StartDate != "" && req.EndDate != "" {
		startDate := req.StartDate + " 00:00:00"
		endDate := req.EndDate + " 23:59:59"
		queryString += " AND payments.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := p.db.Count(
		ctx,
		&model.Payment{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN events ON payments.event_id = events.id
			INNER JOIN users ON events.user_id = users.id
    	`),
	); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var transactions []*model.Payment
	if err := p.db.Find(
		ctx,
		&transactions,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithPreload([]string{"Event", "Event.User"}),
		database.WithJoin(`
			INNER JOIN events ON payments.event_id = events.id
			INNER JOIN users ON events.user_id = users.id
    	`),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return transactions, pagination, nil
}
