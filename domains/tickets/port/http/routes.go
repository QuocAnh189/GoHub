package http

import (
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/tickets/repository"
	"gohub/domains/tickets/service"
	"gohub/internal/libs/validation"
	middleware "gohub/pkg/middleware"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	TicketRepository := repository.NewTicketRepository(sqlDB)
	TicketService := service.NewTicketService(validator, TicketRepository)
	TicketHandler := NewTicketHandler(TicketService)

	authMiddleware := middleware.JWTAuth()

	expenseRoute := r.Group("/tickets").Use(authMiddleware)
	{
		expenseRoute.GET("/get-created-tickets", TicketHandler.GetTicketByCreated)
	}
}
