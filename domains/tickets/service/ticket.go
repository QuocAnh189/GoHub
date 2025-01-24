package service

import (
	"context"
	"gohub/domains/tickets/dto"
	"gohub/domains/tickets/model"
	"gohub/domains/tickets/repository"
	"gohub/internal/libs/validation"
	"gohub/pkg/paging"
)

type ITicketService interface {
	GetCreatedTickets(ctx context.Context, userId string, req *dto.ListTicketReq) ([]*model.Ticket, *paging.Pagination, error)
}

type TicketService struct {
	validator  validation.Validation
	repoTicket repository.ITicketRepository
}

func NewTicketService(validator validation.Validation, repoTicket repository.ITicketRepository) *TicketService {
	return &TicketService{
		validator:  validator,
		repoTicket: repoTicket,
	}
}

func (s *TicketService) GetCreatedTickets(ctx context.Context, userId string, req *dto.ListTicketReq) ([]*model.Ticket, *paging.Pagination, error) {
	tickets, pagination, err := s.repoTicket.GetCreatedTickets(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return tickets, pagination, nil
}
