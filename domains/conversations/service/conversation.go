package service

import (
	"context"
	"gohub/domains/conversations/dto"
	"gohub/domains/conversations/model"
	"gohub/domains/conversations/repository"
	"gohub/internal/libs/logger"
	"gohub/pkg/paging"
	"gohub/pkg/utils"

	"gohub/internal/libs/validation"
)

type IConversationService interface {
	CreateMessage(ctx context.Context, req *dto.CreateMessageReq) (*model.Message, error)
	UpdateMessage(ctx context.Context, id string, req *dto.UpdateMessageReq) (*model.Message, error)
	DeleteMessage(ctx context.Context, messageId string) (*model.Message, error)
	GetConversationsOrganizer(ctx context.Context, organizerId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error)
	GetConversationsByUser(ctx context.Context, userId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error)
	GetMessagesByConversation(ctx context.Context, conversationID string, req *dto.ListMessageReq) ([]*model.Message, *paging.Pagination, error)
}

type ConversationService struct {
	validator        validation.Validation
	repoConversation repository.IConversationRepository
}

func NewConversationService(validator validation.Validation, repo repository.IConversationRepository) *ConversationService {
	return &ConversationService{
		validator:        validator,
		repoConversation: repo,
	}
}

func (c *ConversationService) CreateMessage(ctx context.Context, req *dto.CreateMessageReq) (*model.Message, error) {
	if err := c.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var message model.Message
	utils.MapStruct(&message, req)

	err := c.repoConversation.CreateMessage(ctx, &message)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &message, nil
}

func (c *ConversationService) UpdateMessage(ctx context.Context, id string, req *dto.UpdateMessageReq) (*model.Message, error) {
	if err := c.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	message, err := c.repoConversation.GetMessageById(ctx, id)
	if err != nil {
		logger.Errorf("Update.GetCategoryByID fail, id: %s, error: %s", id, err)
	}

	utils.MapStruct(message, req)
	err = c.repoConversation.UpdateMessage(ctx, message)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return message, nil
}

func (c *ConversationService) DeleteMessage(ctx context.Context, messageId string) (*model.Message, error) {
	err := c.repoConversation.DeleteMessage(ctx, messageId)
	if err != nil {
		return nil, err
	}

	message, err := c.repoConversation.GetMessageDeleteById(ctx, messageId)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (c *ConversationService) GetConversationsOrganizer(ctx context.Context, organizerId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error) {
	reviews, pagination, err := c.repoConversation.GetConversationByOrganizer(ctx, organizerId, req)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (c *ConversationService) GetConversationsByUser(ctx context.Context, userId string, req *dto.ListConversationReq) ([]*model.Conversation, *paging.Pagination, error) {
	reviews, pagination, err := c.repoConversation.GetConversationByUser(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}
	return reviews, pagination, nil
}

func (c *ConversationService) GetMessagesByConversation(ctx context.Context, conversationID string, req *dto.ListMessageReq) ([]*model.Message, *paging.Pagination, error) {
	messages, pagination, err := c.repoConversation.GetMessageByConversation(ctx, conversationID, req)
	if err != nil {
		return nil, nil, err
	}
	return messages, pagination, nil
}
