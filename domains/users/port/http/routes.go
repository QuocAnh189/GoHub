package http

import (
	"gohub/database"
	roleRepository "gohub/domains/roles/repository"
	"gohub/domains/users/repository"
	"gohub/domains/users/service"
	middleware "gohub/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gohub/internal/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	userRepository := repository.NewUserRepository(sqlDB)
	roleRepository := roleRepository.NewRoleRepository(sqlDB)
	userService := service.NewUserService(validator, userRepository, roleRepository)
	userHandler := NewUserHandler(userService)

	authMiddleware := middleware.JWTAuth()
	userRoute := r.Group("/users").Use(authMiddleware)
	{
		userRoute.GET("/", userHandler.GetUsers)
		userRoute.POST("/", userHandler.CreateUser)
		userRoute.GET("/:id", userHandler.GetUserById)
		userRoute.PUT("/:id", userHandler.UpdateUser)
		userRoute.GET("/profile", authMiddleware, userHandler.GetProfile)
		userRoute.PATCH("/change-password", userHandler.ChangePassword)
		userRoute.GET("/:id/followers", userHandler.GetFollowers)
		userRoute.GET("/:id/followings", userHandler.GetFollowing)
		userRoute.PATCH("/follow/:followeeId", userHandler.FollowUser)
		userRoute.PATCH("/unfollow/:followeeId", userHandler.UnfollowUser)
		userRoute.GET("/check-follower/:followeeId", userHandler.CheckFollower)
		userRoute.GET("/invitations", userHandler.GetInvitations)
		userRoute.GET("/check-invitation/:inviteeId", userHandler.CheckInvitation)
		userRoute.PATCH("/invitations", userHandler.InviteUsers)
		userRoute.GET("/notification-following", userHandler.GetNotificationFollowings)
		userRoute.GET("/notification-order", userHandler.GetNotificationOrders)
	}
}
