package http

import (
	"gohub/database"
	"gohub/domains/auth/service"
	"gohub/domains/users/repository"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	userRepository := repository.NewUserRepository(sqlDB)
	AuthService := service.NewAuthService(validator, userRepository)
	authHandler := NewAuthHandler(AuthService)

	authRoute := r.Group("/auth")	
	{
		authRoute.POST("/validate-user", authHandler.ValidateUser)
		authRoute.POST("/signup", authHandler.SignUp)
		authRoute.POST("/signin", authHandler.SignIn)
		authRoute.POST("/signout", authHandler.SignOut)
		authRoute.POST("/external-signin", authHandler.ExternalSignIn)
		authRoute.GET("/external-callback", authHandler.ExternalCallback)
		authRoute.POST("/refresh-token", authHandler.RefreshToken)
		authRoute.POST("/forgot-password", authHandler.ForgotPassword)
		authRoute.POST("/reset-password", authHandler.ResetPassword)
		authRoute.GET("/profile", authHandler.GetProfile)
	}

}