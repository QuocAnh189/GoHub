package http

import (
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/categories/repository"
	"gohub/domains/categories/service"
	"gohub/internal/libs/validation"
	middleware "gohub/pkg/middleware"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	CategoryRepository := repository.NewCategoryRepository(sqlDB)
	CategoryService := service.NewCategoryService(validator, CategoryRepository)
	CategoryHandler := NewCategoryHandler(CategoryService)

	authMiddleware := middleware.JWTAuth()

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.GET("/", CategoryHandler.GetCategories)
		categoryRoute.POST("/", authMiddleware, CategoryHandler.CreateCategory)
		categoryRoute.GET("/:id", authMiddleware, CategoryHandler.GetCategoryById)
		categoryRoute.PUT("/:id", authMiddleware, CategoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:id", authMiddleware, CategoryHandler.DeleteCategory)
		categoryRoute.DELETE("/", authMiddleware, CategoryHandler.DeleteMultipleCategory)
		categoryRoute.PATCH("/restore", authMiddleware, CategoryHandler.RestoreCategories)
	}
}
