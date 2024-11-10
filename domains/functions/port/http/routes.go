package http

import (
	"gohub/database"
	"gohub/domains/functions/repository"
	"gohub/domains/functions/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	functionRepository := repository.NewFunctionRepository(sqlDB)
	functionService := service.NewFunctionService(validator, functionRepository)
	functionHandler := NewFunctionHandler(functionService)

	functionRoute := r.Group("/functions")
	{
		functionRoute.GET("/", functionHandler.GetFunctions)
		functionRoute.POST("/", functionHandler.CreateFunction)
		functionRoute.GET("/:id", functionHandler.GetFunction)
		functionRoute.PUT("/:id", functionHandler.UpdateFunction)
		functionRoute.DELETE("/:id", functionHandler.DeleteFunction)
		functionRoute.POST("/:id/enable-command/:commandId", functionHandler.EnableCommand)
		functionRoute.POST("/:id/disable-command/:commandId", functionHandler.DisableCommand)
	}
}