package http

import (
	"gohub/database"
	"gohub/domains/commands/repository"
	"gohub/domains/commands/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
    commandRepository := repository.NewCommandRepository(sqlDB)
	commandService := service.NewCommandService(validator, commandRepository)
	commandHandler := NewCommandHandler(commandService)

	commandRoute := r.Group("/commands")
	{
		commandRoute.GET("/get-in-function/:functionId", commandHandler.GetInFunction)
		commandRoute.GET("/get-not-in-function/:functionId", commandHandler.GetNotInFunction)
	}
}