package http

import (
	"gohub/database"
	"gohub/domains/conversations/repository"
	"gohub/domains/conversations/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	conversationRepository := repository.NewConversationRepository(sqlDB)
	conversationService := service.NewConversationService(validator, conversationRepository)
	conversationHandler := NewConversationHandler(conversationService)

	conversationRoute := r.Group("/conversations")
	{
		conversationRoute.GET("/get-by-event/:eventId", conversationHandler.GetConversationsByEvent)
		conversationRoute.GET("/get-by-user/:userId", conversationHandler.GetConversationsByUser)
		conversationRoute.GET("/:id/messages", conversationHandler.GetConversationsByEvent)
	}

}