package service

import (
	"context"
	"gohub/domains/payments/dto"
	"gohub/domains/payments/model"
	"gohub/domains/payments/repository"
	"gohub/internal/libs/validation"
	"gohub/pkg/paging"
)

type IPaymentService interface {
	GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error)
	GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error)
}

type PaymentService struct {
	validator   validation.Validation
	repoPayment repository.IPaymentRepository
}

func NewPaymentService(validator validation.Validation, repoPayment repository.IPaymentRepository) *PaymentService {
	return &PaymentService{
		validator:   validator,
		repoPayment: repoPayment,
	}
}

func (s *PaymentService) GetTransactions(ctx context.Context, userId string, req *dto.ListTransactionReq) ([]*model.Payment, *paging.Pagination, error) {
	transactions, pagination, err := s.repoPayment.GetTransactions(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return transactions, pagination, nil
}

func (s *PaymentService) GetOrders(ctx context.Context, userId string, req *dto.ListOrderReq) ([]*model.Payment, *paging.Pagination, error) {
	orders, pagination, err := s.repoPayment.GetOrders(ctx, userId, req)
	if err != nil {
		return nil, nil, err
	}

	return orders, pagination, nil
}
