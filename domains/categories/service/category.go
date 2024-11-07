package service

import (
	"context"
	"gohub/domains/categories/model"
	"gohub/domains/categories/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type ICategoryService interface {
	CreateCategory(ctx context.Context)	error
	GetCategory(ctx context.Context) (*model.Category, error)
	GetCategories(ctx context.Context)	([]*model.Category, error)
	UpdateCategory(ctx context.Context) error
	DeleteCategory(ctx context.Context) error
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

//	@Summary	 Create a new category
//  @Description Creates a new category based on the provided details. The request must include multipart form data.
//	@Tags		 Categories
//	@Produce	 json
//	@Success	 201	{object}	response.Response	"Category created successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/categories [post]
func (s *CategoryService) CreateCategory(ctx context.Context) error {
	panic("unimplemented")
}

//	@Summary	 Retrieve a list of categories
//  @Description Fetches a paginated list of categories based on the provided filter parameters.
//	@Tags		 Categories
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Category created successfully"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/categories [get]
func (s *CategoryService) GetCategories(ctx context.Context) ([]*model.Category, error) {
	panic("unimplemented")
}

//	@Summary	 Retrieve a category by its ID
//  @Description Fetches the details of a specific category based on the provided category ID.
//	@Tags		 Categories
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Category created successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/categories/{categoryId} [get]
func (s *CategoryService) GetCategory(ctx context.Context) (*model.Category, error) {
	panic("unimplemented")
}

//	@Summary	 Update an existing category
//  @Description Updates the details of an existing category based on the provided category ID and update information.
//	@Tags		 Categories
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Category updated successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/categories/{categoryId} [put]
func (s *CategoryService) UpdateCategory(ctx context.Context) error {
	panic("unimplemented")
}


//	@Summary	 Delete a category
//  @Description Deletes the category with the specified ID.
//	@Tags		 Categories
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Category updated successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/categories/{categoryId} [delete]
func (s *CategoryService) DeleteCategory(ctx context.Context) error {
	panic("unimplemented")
}
