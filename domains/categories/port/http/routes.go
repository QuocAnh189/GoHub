package http

import (
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/categories/repository"
	"gohub/domains/categories/service"
	middleware "gohub/pkg/middlewares"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	CategoryRepository := repository.NewCategoryRepository(sqlDB)
	CategoryService := service.NewCategoryService(validator, CategoryRepository)
	CategoryHandler := NewCategoryHandler(CategoryService)

	authMiddleware := middleware.JWTAuth()

	categoryRoute := r.Group("/categories").Use(authMiddleware)
	{
		categoryRoute.GET("/", CategoryHandler.GetCategories)
		categoryRoute.POST("/", CategoryHandler.CreateCategory)
		categoryRoute.GET("/:id", CategoryHandler.GetCategoryById)
		categoryRoute.PUT("/:id", CategoryHandler.UpdateCategory)
		categoryRoute.DELETE("/", CategoryHandler.DeleteMultipleCategory)
		categoryRoute.DELETE("/:id", CategoryHandler.DeleteCategory)
		categoryRoute.PATCH("/restore", CategoryHandler.RestoreCategories)
	}
}
