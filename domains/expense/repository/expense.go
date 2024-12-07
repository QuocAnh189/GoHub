package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/expense/dto"
	"gohub/domains/expense/model"
	"gohub/pkg/paging"
	"gohub/pkg/utils"
)

type IExpenseRepository interface {
	GetExpenseById(ctx context.Context, id string) (*model.Expense, error)
	GetExpensesByEventId(ctx context.Context, eventId string, req *dto.ListExpenseReq) ([]*model.Expense, *paging.Pagination, error)
	Create(ctx context.Context, expense *model.Expense) error
	Update(ctx context.Context, expense *model.Expense) error
	Delete(ctx context.Context, id string) error
	GetSubExpenseById(ctx context.Context, subExpenseId string) (*model.SubExpense, error)
	CreateSubExpense(ctx context.Context, subExpense *model.SubExpense) error
	UpdateSubExpense(ctx context.Context, subExpenseId string, req *dto.UpdateSubExpenseReq) error
	DeleteSubExpense(ctx context.Context, subExpenseId string) error
}

type ExpenseRepository struct {
	db database.IDatabase
}

func NewExpenseRepository(db database.IDatabase) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (e *ExpenseRepository) GetExpenseById(ctx context.Context, id string) (*model.Expense, error) {
	var expense model.Expense
	if err := e.db.FindById(ctx, id, &expense); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (e *ExpenseRepository) GetExpensesByEventId(ctx context.Context, eventId string, req *dto.ListExpenseReq) ([]*model.Expense, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Name != "" {
		query = append(query, database.NewQuery("event_id = ? AND name LIKE ?", eventId, "%"+req.Name+"%"))
	} else {
		query = append(query, database.NewQuery("event_id = ? ", eventId))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := e.db.Count(ctx, &model.Expense{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var expenses []*model.Expense
	if err := e.db.Find(
		ctx,
		&expenses,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
		database.WithPreload([]string{"SubExpenses"}),
	); err != nil {
		return nil, nil, err
	}

	return expenses, pagination, nil
}

func (e *ExpenseRepository) Create(ctx context.Context, expense *model.Expense) error {
	return e.db.Create(ctx, expense)
}

func (e *ExpenseRepository) Update(ctx context.Context, expense *model.Expense) error {
	return e.db.Update(ctx, expense)
}

func (e *ExpenseRepository) Delete(ctx context.Context, id string) error {
	expense, err := e.GetExpenseById(ctx, id)
	if err != nil {
		return err
	}
	return e.db.Delete(ctx, expense)
}

func (e *ExpenseRepository) GetSubExpenseById(ctx context.Context, id string) (*model.SubExpense, error) {
	var expense model.SubExpense
	if err := e.db.FindById(ctx, id, &expense); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (e *ExpenseRepository) CreateSubExpense(ctx context.Context, subExpense *model.SubExpense) error {
	handler := func() error {
		if subExpense.Price > 0 {
			var expense model.Expense
			if err := e.db.FindById(ctx, subExpense.ExpenseId, &expense); err != nil {
				return err
			}
			expense.Total += subExpense.Price
			if err := e.db.Update(ctx, expense); err != nil {
				return err
			}
		}

		if err := e.db.Create(ctx, subExpense); err != nil {
			return err
		}
		return nil
	}

	if err := e.db.WithTransaction(handler); err != nil {
		return err
	}
	return nil
}

func (e *ExpenseRepository) UpdateSubExpense(ctx context.Context, subExpenseId string, req *dto.UpdateSubExpenseReq) error {
	handler := func() error {
		subExpense, err := e.GetSubExpenseById(ctx, subExpenseId)
		if err != nil {
			return err
		}

		if subExpense.Price != req.Price {
			var expense model.Expense
			if err := e.db.FindById(ctx, subExpense.ExpenseId, &expense); err != nil {
				return err
			}
			expense.Total = expense.Total - subExpense.Price + req.Price
			if err := e.db.Update(ctx, expense); err != nil {
				return err
			}
		}

		utils.MapStruct(subExpense, req)
		if err := e.db.Update(ctx, subExpense); err != nil {
			return err
		}
		return nil
	}

	if err := e.db.WithTransaction(handler); err != nil {
		return err
	}
	return nil
}

func (e *ExpenseRepository) DeleteSubExpense(ctx context.Context, subExpenseId string) error {
	handler := func() error {
		subExpense, err := e.GetSubExpenseById(ctx, subExpenseId)
		if err != nil {
			return err
		}

		if subExpense.Price != 0 {
			var expense model.Expense
			if err := e.db.FindById(ctx, subExpense.ExpenseId, &expense); err != nil {
				return err
			}
			expense.Total -= subExpense.Price
			if err := e.db.Update(ctx, expense); err != nil {
				return err
			}
			return nil
		}

		if err := e.db.Delete(ctx, subExpense); err != nil {
			return err
		}
		return nil
	}

	if err := e.db.WithTransaction(handler); err != nil {
		return err
	}
	return nil
}
