package repository

import (
	"context"
	"gohub/configs"
	"gohub/database"
	"gohub/domains/categories/dto"
	"gohub/domains/categories/model"
	"gohub/pkg/paging"
)

type ICategoryRepository interface {
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
	DeleteByIds(ctx context.Context, ids []string) error
	RestoreByIds(ctx context.Context, ids []string) error
	ListCategories(ctx context.Context, req *dto.ListCategoryReq) ([]*model.Category, *paging.Pagination, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*model.Category, error)
}

type CategoryRepository struct {
	db database.IDatabase
}

func NewCategoryRepository(db database.IDatabase) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c *CategoryRepository) ListCategories(ctx context.Context, req *dto.ListCategoryReq) ([]*model.Category, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Name != "" {
		query = append(query, database.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := c.db.Count(ctx, &model.Category{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	if req.TakeAll {
		pagination.PageSize = total
	}

	var categories []*model.Category
	if err := c.db.Find(
		ctx,
		&categories,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.PageSize)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (c *CategoryRepository) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	var category model.Category
	if err := c.db.FindById(ctx, id, &category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	var category model.Category
	query := database.NewQuery("name = ?", name)
	if err := c.db.FindOne(ctx, &category, database.WithQuery(query)); err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepository) Create(ctx context.Context, category *model.Category) error {
	return c.db.Create(ctx, category)
}

func (c *CategoryRepository) Update(ctx context.Context, category *model.Category) error {
	return c.db.Update(ctx, category)
}

func (c *CategoryRepository) Delete(ctx context.Context, id string) error {
	category, err := c.GetCategoryById(ctx, id)
	if err != nil {
		return err
	}
	return c.db.Delete(ctx, category)
}

func (c *CategoryRepository) DeleteByIds(ctx context.Context, ids []string) error {
	return c.db.DeleteByIds(ctx, &model.Category{}, ids)
}

func (c *CategoryRepository) RestoreByIds(ctx context.Context, ids []string) error {
	return c.db.RestoreByIds(ctx, &model.Category{}, ids)
}
