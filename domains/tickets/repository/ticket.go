package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/tickets/dto"
	"gohub/domains/tickets/model"
	"gohub/pkg/paging"
)

type ITicketRepository interface {
	GetCreatedTickets(ctx context.Context, userId string, req *dto.ListTicketReq) ([]*model.Ticket, *paging.Pagination, error)
}

type TicketRepository struct {
	db database.IDatabase
}

func NewTicketRepository(db database.IDatabase) *TicketRepository {
	return &TicketRepository{db: db}
}

func (t *TicketRepository) GetCreatedTickets(ctx context.Context, userId string, req *dto.ListTicketReq) ([]*model.Ticket, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "tickets.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND (customer_name ILIKE ? OR customer_email ILIKE ? OR customer_phone ILIKE ? OR events.name ILIKE ? OR ticket_types.name ILIKE ?)"
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
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
	if err := t.db.Count(
		ctx,
		&model.Ticket{},
		&total,
		database.WithJoin(`
			INNER JOIN events ON tickets.event_id = events.id
			INNER JOIN ticket_types ON tickets.ticket_type_id = ticket_types.id
    	`),
		database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var tickets []*model.Ticket
	if err := t.db.Find(
		ctx,
		&tickets,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithJoin(`
			INNER JOIN events ON tickets.event_id = events.id
			INNER JOIN ticket_types ON tickets.ticket_type_id = ticket_types.id
    	`),
		database.WithPreload([]string{"Event", "TicketType"}),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return tickets, pagination, nil
}
