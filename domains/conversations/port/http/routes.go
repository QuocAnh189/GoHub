package http

import (
	"gohub/database"
	"gohub/domains/conversations/repository"
	"gohub/domains/conversations/service"
	middleware "gohub/pkg/middlewares"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	conversationRepository := repository.NewConversationRepository(sqlDB)
	conversationService := service.NewConversationService(validator, conversationRepository)
	conversationHandler := NewConversationHandler(conversationService)

	authMiddleware := middleware.JWTAuth()
	conversationRoute := r.Group("/conversations").Use(authMiddleware)
	{
		conversationRoute.GET("/get-by-organizer/:organizerId", conversationHandler.GetConversationsByOrganizer)
		conversationRoute.GET("/get-by-user/:userId", conversationHandler.GetConversationsByUser)
		conversationRoute.GET("/:id/messages", conversationHandler.GetMessagesByConversation)
		conversationRoute.POST("/:id/messages", conversationHandler.CreateMessage)
		conversationRoute.PUT("/:id/messages/:messageId", conversationHandler.UpdateMessage)
		conversationRoute.DELETE("/:id/messages/:messageId", conversationHandler.DeleteMessage)
	}
}
