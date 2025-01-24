package http

import (
	"gohub/database"
	"gohub/domains/events/repository"
	"gohub/domains/events/service"
	middleware "gohub/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gohub/internal/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	eventRepository := repository.NewEventRepository(sqlDB)
	eventService := service.NewEventService(validator, eventRepository)
	eventHandler := NewEventHandler(eventService)

	authMiddleware := middleware.JWTAuth()
	eventRoute := r.Group("/events")
	{
		eventRoute.GET("/", eventHandler.GetEvents)
		eventRoute.POST("/", authMiddleware, eventHandler.CreateEvent)
		eventRoute.GET("/:id", eventHandler.GetEvent)
		eventRoute.PUT("/:id", authMiddleware, eventHandler.UpdateEvent)
		eventRoute.DELETE("/:id", authMiddleware, eventHandler.DeleteEvent)
		eventRoute.DELETE("/", authMiddleware, eventHandler.DeleteMultipleEvent)
		eventRoute.GET("/get-created-events", authMiddleware, eventHandler.GetCreatedEvent)
		eventRoute.GET("/get-created-events-analysis", authMiddleware, eventHandler.GetCreatedEventAnalysis)
		eventRoute.PATCH("/restore", authMiddleware, eventHandler.RestoreEvents)
		eventRoute.GET("/get-deleted-events", authMiddleware, eventHandler.GetTrashedEvent)
		eventRoute.PATCH("/favourite/:id", authMiddleware, eventHandler.FavouriteEvent)
		eventRoute.PATCH("/unfavourite/:id", authMiddleware, eventHandler.UnFavouriteEvent)
		eventRoute.GET("/get-favourite-events", authMiddleware, eventHandler.GetFavouriteEvent)
		eventRoute.PATCH("/make-events-private", authMiddleware, eventHandler.MakeEventPrivate)
		eventRoute.PATCH("/make-events-public", authMiddleware, eventHandler.MakeEventPublic)
		eventRoute.PATCH("/apply-coupons/:id", authMiddleware, eventHandler.ApplyCoupons)
		eventRoute.GET("/check-favourite/:id", authMiddleware, eventHandler.CheckFavourite)
	}
}
