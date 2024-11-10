package http

import (
	"gohub/database"
	"gohub/domains/roles/repository"
	"gohub/domains/roles/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation){
	roleRepository := repository.NewRoleRepository(sqlDB)
	roleService := service.NewRolService(validator, roleRepository)
	roleHandler := NewRoleHandler(roleService)

	roleRoute := r.Group("/roles")
	{
		roleRoute.POST("/:{id}/add-function/{functionId}", roleHandler.AddFunction)
		roleRoute.POST("/:{id}/remove-function/{functionId}", roleHandler.RemoveFunction)
	}
}