package service

import (
	"context"
	"errors"
	"gohub/domains/categories/dto"
	"gohub/domains/categories/model"
	"gohub/domains/categories/repository"
	"gohub/internal/libs/logger"
	"gohub/pkg/messages"
	"gohub/pkg/paging"
	"gohub/pkg/utils"

	"gohub/internal/libs/validation"
)

type ICategoryService interface {
	CreateCategory(ctx context.Context, req *dto.CreateCategoryReq) (*model.Category, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategories(ctx context.Context, req *dto.ListCategoryReq) ([]*model.Category, *paging.Pagination, error)
	UpdateCategory(ctx context.Context, id string, req *dto.UpdateCategoryReq) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	DeleteCategories(ctx context.Context, ids *dto.DeleteRequest) error
	RestoreCategories(ctx context.Context, ids *dto.RestoreRequest) error
}

type CategoryService struct {
	validator validation.Validation
	repo      repository.ICategoryRepository
}

func NewCategoryService(validator validation.Validation, repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		validator: validator,
		repo:      repo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *dto.CreateCategoryReq) (*model.Category, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	_, err := s.repo.GetCategoryByName(ctx, req.Name)
	if err == nil {
		return nil, errors.New(messages.CategoryNameExists)
	}

	var category model.Category
	utils.MapStruct(&category, req)

	err = s.repo.Create(ctx, &category)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) GetCategories(ctx context.Context, req *dto.ListCategoryReq) ([]*model.Category, *paging.Pagination, error) {
	categories, pagination, err := s.repo.ListCategories(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (s *CategoryService) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	category, err := s.repo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id string, req *dto.UpdateCategoryReq) (*model.Category, error) {
	if err := s.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	category, err := s.repo.GetCategoryById(ctx, id)
	if err != nil {
		logger.Errorf("Update.GetCategoryByID fail, id: %s, error: %s", id, err)
		return nil, errors.New(messages.CategoryNotFound)
	}

	categoryExistsName, err := s.repo.GetCategoryByName(ctx, req.Name)
	if err == nil && categoryExistsName.Name != category.Name {
		return nil, errors.New(messages.CategoryNameExists)
	}

	utils.MapStruct(category, req)
	err = s.repo.Update(ctx, category)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteCategories(ctx context.Context, ids *dto.DeleteRequest) error {
	var err error
	if len(ids.Ids) == 1 {
		err = s.repo.Delete(ctx, ids.Ids[0])
	} else {
		err = s.repo.DeleteByIds(ctx, ids.Ids)
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) RestoreCategories(ctx context.Context, ids *dto.RestoreRequest) error {
	err := s.repo.RestoreByIds(ctx, ids.Ids)

	if err != nil {
		return err
	}

	return nil
}
