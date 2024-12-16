package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/conversations/dto"
	"gohub/domains/conversations/model"
	"gohub/pkg/paging"
)

type IConversationRepository interface {
	CreateMessage(ctx context.Context, message *model.Message) error
	UpdateMessage(ctx context.Context, message *model.Message) error
	DeleteMessage(ctx context.Context, messageId string) error
	GetConversationByOrganizer(ctx context.Context, organizerId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error)
	GetConversationByUser(ctx context.Context, userId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error)
	GetMessageByConversation(ctx context.Context, conservationId string, req *dto.ListMessageReq) ([]*model.Message, *paging.Pagination, error)
	GetMessageById(ctx context.Context, messageId string) (*model.Message, error)
	GetMessageDeleteById(ctx context.Context, messageId string) (*model.Message, error)
}

type ConversationRepo struct {
	db database.IDatabase
}

func NewConversationRepository(db database.IDatabase) *ConversationRepo {
	return &ConversationRepo{db: db}
}

func (c *ConversationRepo) CreateMessage(ctx context.Context, message *model.Message) error {
	return c.db.Create(ctx, message)
}

func (c *ConversationRepo) UpdateMessage(ctx context.Context, message *model.Message) error {
	return c.db.Update(ctx, message)
}

func (c *ConversationRepo) DeleteMessage(ctx context.Context, messageId string) error {
	message, err := c.GetMessageById(ctx, messageId)
	if err != nil {
		return err
	}

	return c.db.Delete(ctx, message)
}

func (c *ConversationRepo) GetConversationByOrganizer(ctx context.Context, organizerId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "organizer_id = ?"
	args = append(args, organizerId)

	if req.Search != "" {
		queryString += " AND users.user_name LIKE ? OR users.full_name LIKE ?"
		args = append(args, "%"+req.Search+"%", "%"+req.Search+"%")
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	query = append(query, database.NewQuery(queryString, args...))

	var total int64
	if err := c.db.Count(
		ctx,
		&model.Conversation{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN users ON conversations.user_id = users.id
    	`),
	); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var conversations []*model.Conversation
	if err := c.db.Find(
		ctx,
		&conversations,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithJoin(`
			INNER JOIN users ON conversations.user_id = users.id
    	`),
		database.WithPreload([]string{"User", "LastMessage", "Event"}),
	); err != nil {
		return nil, nil, err
	}

	return conversations, pagination, nil
}

func (c *ConversationRepo) GetConversationByUser(ctx context.Context, userId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	//query := make([]database.Query, 0)
	//query = append(query, database.NewQuery("user_id = ?", userId))

	query := make([]database.Query, 0)
	args := make([]interface{}, 0)

	queryString := "conversations.user_id = ?"
	args = append(args, userId)

	if req.Search != "" {
		queryString += " AND events.name LIKE ?"
		args = append(args, "%"+req.Search+"%")
	}

	query = append(query, database.NewQuery(queryString, args...))

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := c.db.Count(
		ctx,
		&model.Conversation{},
		&total,
		database.WithQuery(query...),
		database.WithJoin(`
			INNER JOIN events ON conversations.event_id = events.id
    	`),
	); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var conversations []*model.Conversation
	if err := c.db.Find(
		ctx,
		&conversations,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithJoin(`
			INNER JOIN events ON conversations.event_id = events.id
    	`),
		database.WithPreload([]string{"Organizer", "LastMessage", "Event"}),
	); err != nil {
		return nil, nil, err
	}

	return conversations, pagination, nil
}

func (c *ConversationRepo) GetMessageByConversation(ctx context.Context, conservationId string, req *dto.ListMessageReq) ([]*model.Message, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	query = append(query, database.NewQuery("conversation_id = ?", conservationId))

	order := "created_at"

	var total int64
	if err := c.db.Count(ctx, &model.Message{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var messages []*model.Message
	if err := c.db.Find(
		ctx,
		&messages,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return messages, pagination, nil
}

func (c *ConversationRepo) GetMessageById(ctx context.Context, messageId string) (*model.Message, error) {
	var message model.Message
	if err := c.db.FindById(ctx, messageId, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (c *ConversationRepo) GetMessageDeleteById(ctx context.Context, messageId string) (*model.Message, error) {
	var message model.Message
	if err := c.db.FindDeleteById(ctx, messageId, &message); err != nil {
		return nil, err
	}

	return &message, nil
}
