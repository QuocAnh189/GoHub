package service

import (
	"context"
	"errors"
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/jackc/pgx/v5/pgconn"
	"gohub/domains/expense/dto"
	"gohub/domains/expense/model"
	"gohub/domains/expense/repository"
	"gohub/pkg/messages"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
)

type IExpenseService interface {
	GetExpenseByEvent(ctx context.Context, userId string, req *dto.ListExpenseReq) ([]*model.Expense, *paging.Pagination, error)
	GetExpenseById(ctx context.Context, id string) (*model.Expense, error)
	CreateExpense(ctx context.Context, req *dto.CreatedExpenseReq) (*model.Expense, error)
	UpdateExpense(ctx context.Context, id string, req *dto.UpdatedExpenseReq) (*model.Expense, error)
	DeleteExpense(ctx context.Context, id string) error
	CreateSubExpense(ctx context.Context, req *dto.CreatedSubExpenseReq) (*model.SubExpense, error)
	UpdateSubExpense(ctx context.Context, subExpenseId string, req *dto.UpdateSubExpenseReq) error
	DeleteSubExpense(ctx context.Context, subExpenseId string) error
}

type ExpenseService struct {
	validator   validation.Validation
	repoExpense repository.IExpenseRepository
}

func NewExpenseService(validator validation.Validation, repoExpense repository.IExpenseRepository) *ExpenseService {
	return &ExpenseService{
		validator:   validator,
		repoExpense: repoExpense,
	}
}

func (e *ExpenseService) GetExpenseByEvent(ctx context.Context, eventId string, req *dto.ListExpenseReq) ([]*model.Expense, *paging.Pagination, error) {
	expenses, pagination, err := e.repoExpense.GetExpensesByEventId(ctx, eventId, req)
	if err != nil {
		return nil, nil, err
	}

	return expenses, pagination, nil
}

func (e *ExpenseService) GetExpenseById(ctx context.Context, id string) (*model.Expense, error) {
	expense, err := e.repoExpense.GetExpenseById(ctx, id)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *ExpenseService) CreateExpense(ctx context.Context, req *dto.CreatedExpenseReq) (*model.Expense, error) {
	if err := e.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var expense model.Expense
	utils.MapStruct(&expense, req)

	err := e.repoExpense.Create(ctx, &expense)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, errors.New(messages.TitleExpenseAlreadyExists)
		}
		return nil, err
	}

	return &expense, nil
}

func (e *ExpenseService) UpdateExpense(ctx context.Context, id string, req *dto.UpdatedExpenseReq) (*model.Expense, error) {
	if err := e.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	expense, err := e.repoExpense.GetExpenseById(ctx, id)
	if err != nil {
		logger.Errorf("Update.GetCategoryByID fail, id: %s, error: %s", id, err)
		return nil, errors.New(messages.ExpenseNotFound)
	}

	utils.MapStruct(expense, req)
	err = e.repoExpense.Update(ctx, expense)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return expense, nil
}

func (e *ExpenseService) DeleteExpense(ctx context.Context, id string) error {
	err := e.repoExpense.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (e *ExpenseService) CreateSubExpense(ctx context.Context, req *dto.CreatedSubExpenseReq) (*model.SubExpense, error) {
	if err := e.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	var subExpense model.SubExpense
	utils.MapStruct(&subExpense, req)

	err := e.repoExpense.CreateSubExpense(ctx, &subExpense)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, errors.New(messages.TitleExpenseAlreadyExists)
		}
		return nil, err
	}

	return &subExpense, nil
}

func (e *ExpenseService) UpdateSubExpense(ctx context.Context, subExpenseId string, req *dto.UpdateSubExpenseReq) error {
	if err := e.validator.ValidateStruct(req); err != nil {
		return err
	}

	err := e.repoExpense.UpdateSubExpense(ctx, subExpenseId, req)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", subExpenseId, err)
		return err
	}

	return nil
}

func (e *ExpenseService) DeleteSubExpense(ctx context.Context, subExpenseId string) error {
	if err := e.repoExpense.DeleteSubExpense(ctx, subExpenseId); err != nil {
		return err
	}

	return nil
}
