package http

import (
	"gohub/domains/categories/service"
	"gohub/pkg/response"
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

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateCategory route"})
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetCategories route"})
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetCategory route"})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateCategory route"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {	
	response.JSON(c, http.StatusOK, gin.H{"message": "Test DeleteCategory route"})
}