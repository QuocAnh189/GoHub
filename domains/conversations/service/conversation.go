package service

import (
	"context"
	"gohub/domains/conversations/model"
	"gohub/domains/conversations/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IConversationService interface {
	GetConversationsByEvent(ctx context.Context, eventID string) ([]*model.Conversation, error)
	GetConversationsByUser(ctx context.Context, userId string) ([]*model.Conversation, error)
	GetMessagesByConversation(ctx context.Context, conversationID string) ([]*model.Message, error)
}


type ConversationService struct {
	validator validation.Validation
	repo      repository.IConversationRepository
}

func NewConversationService(validator validation.Validation, repo repository.IConversationRepository) *ConversationService {
	return &ConversationService{
		validator: validator,
		repo:      repo,
	}
}

func (c *ConversationService) GetConversationsByEvent(ctx context.Context, eventID string) ([]*model.Conversation, error) {
	panic("unimplemented")
}

func (c *ConversationService) GetConversationsByUser(ctx context.Context, userId string) ([]*model.Conversation, error) {
	panic("unimplemented")
}

func (c *ConversationService) GetMessagesByConversation(ctx context.Context, conversationID string) ([]*model.Message, error) {
	panic("unimplemented")
}