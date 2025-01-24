package http

import (
	"gohub/database"
	"gohub/domains/auth/service"
	roleRepository "gohub/domains/roles/repository"
	"gohub/domains/users/repository"
	middleware "gohub/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gohub/internal/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	userRepository := repository.NewUserRepository(sqlDB)
	roleRepository := roleRepository.NewRoleRepository(sqlDB)
	AuthService := service.NewAuthService(validator, userRepository, roleRepository)
	authHandler := NewAuthHandler(AuthService)

	authMiddleware := middleware.JWTAuth()
	refreshAuthMiddleware := middleware.JWTRefresh()
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/validate-user", authHandler.ValidateUser)
		authRoute.POST("/signup", authHandler.SignUp)
		authRoute.POST("/signin", authHandler.SignIn)
		authRoute.POST("/external-login", authHandler.ExternalSignIn)
		authRoute.GET("/external-auth-callback", authHandler.ExternalCallback)
		authRoute.POST("/signout", authMiddleware, authHandler.SignOut)
		authRoute.POST("/refresh-token", refreshAuthMiddleware, authHandler.RefreshToken)
		authRoute.POST("/forgot-password", authHandler.ForgotPassword)
		authRoute.POST("/reset-password", authMiddleware, authHandler.ResetPassword)
	}
}
