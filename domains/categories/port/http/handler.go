package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"gohub/domains/categories/dto"
	"gohub/domains/categories/service"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

//		@Summary	 Create a new category
//	 @Description Creates a new category based on the provided details. The request must include multipart form data.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	category, err := h.service.CreateCategory(c, &req)
	if err != nil {
		logger.Error("Failed to get body", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Failed to create category")
		return
	}

	var res dto.Category
	utils.MapStruct(&res, &category)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a list of categories
//	 @Description Fetches a paginated list of categories based on the provided filter parameters.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	var req dto.ListCategoryReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListCategoryRes
	categories, pagination, err := h.service.GetCategories(c, &req)
	if err != nil {
		logger.Error("Failed to get list products: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res.Category, &categories)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)

}

//		@Summary	 Retrieve a category by its ID
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories/{categoryId} [get]
func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	var res dto.Category

	categoryId := c.Param("id")
	category, err := h.service.GetCategoryById(c, categoryId)
	if err != nil {
		logger.Error("Failed to get category detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res, &category)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing category
//	 @Description Updates the details of an existing category based on the provided category ID and update information.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories/{categoryId} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	productId := c.Param("id")
	var req dto.UpdateCategoryReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	category, err := h.service.UpdateCategory(c, productId, &req)
	if err != nil {
		logger.Error("Failed to get body", err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Failed to create category")
		return
	}

	var res dto.Category
	utils.MapStruct(&res, &category)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete a category
//	 @Description Deletes the category with the specified ID.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories/{categoryId} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")

	err := h.service.DeleteCategory(c, categoryId)

	if err != nil {
		logger.Error("Failed to delete category: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete category successfully")
}

//		@Summary	 Delete a categories
//	 @Description Deletes the category with the specified ID.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories [delete]
func (h *CategoryHandler) DeleteCategories(c *gin.Context) {
	var req dto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.DeleteCategories(c, &req)

	if err != nil {
		logger.Error("Failed to delete category: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete category successfully")
}

//		@Summary	 Restore a categories
//	 @Description Deletes the category with the specified ID.
//		@Tags		 Categories
//		@Produce	 json
//		@Success	 200	string		response.Response	"Restore category successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/categories/restore [patch]
func (h *CategoryHandler) RestoreCategories(c *gin.Context) {
	var req dto.RestoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.RestoreCategories(c, &req)

	if err != nil {
		logger.Error("Failed to restore category: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Restore category successfully")
}
