package http

import (
	"gohub/database"
	"gohub/domains/users/repository"
	"gohub/domains/users/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	userRepository := repository.NewUserRepository(sqlDB)
	userService := service.NewUserService(validator, userRepository)
	userHandler := NewUserHandler(userService)

	userRoute := r.Group("/users")
	{
		userRoute.GET("/", userHandler.GetUsers)
		userRoute.POST("/", userHandler.CreateUser)
		userRoute.GET("/:id", userHandler.GetUser)
		userRoute.PUT("/:id", userHandler.UpdateUser)
		userRoute.PATCH("/:id/change-password", userHandler.ChangePassword)
		userRoute.GET("/:id/followers", userHandler.GetFollowers)
		userRoute.GET("/:id/following-users", userHandler.GetFollowing)
		userRoute.PATCH("/follow/:followedUserId", userHandler.FollowUser)
		userRoute.PATCH("/unfollow/:followedUserId", userHandler.UnfollowUser)
	}
}