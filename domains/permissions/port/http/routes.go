package http

import (
	"gohub/database"
	"gohub/domains/permissions/repository"
	"gohub/domains/permissions/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {

	userRepository := repository.NewPermissionRepository(sqlDB)
	userService := service.NewPermissionService(validator, userRepository)
	userHandler := NewPermissionHandler(userService)
	permissionRoute := r.Group("/permissions")
	{
		permissionRoute.GET("/", userHandler.GetPermissions)
		permissionRoute.GET("/roles", userHandler.GetPermissionsByRoles)
		permissionRoute.GET("/get-by-user/:userId", userHandler.GetPermissionsByUsers)
	}
}