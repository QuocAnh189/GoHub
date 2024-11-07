package http

import (
	"gohub/database"
	"gohub/domains/categories/repository"
	"gohub/domains/categories/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	CategoryRepository := repository.NewCategoryRepository(sqlDB)
	CategoryService := service.NewCategoryService(validator, CategoryRepository)
	CategoryHandler := NewCategoryHandler(CategoryService)

	categoryRoute := r.Group("/categories")
	{
		categoryRoute.GET("/", CategoryHandler.GetCategories)
		categoryRoute.POST("/", CategoryHandler.CreateCategory)
		categoryRoute.GET("/:id", CategoryHandler.GetCategory)
		categoryRoute.PUT("/:id", CategoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:id", CategoryHandler.DeleteCategory)
	}

}