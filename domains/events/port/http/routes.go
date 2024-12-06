package http

import (
	"gohub/database"
	"gohub/domains/events/repository"
	"gohub/domains/events/service"
	middleware "gohub/pkg/middlewares"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	eventRepository := repository.NewEventRepository(sqlDB)
	eventService := service.NewEventService(validator, eventRepository)
	eventHandler := NewEventHandler(eventService)

	authMiddleware := middleware.JWTAuth()
	eventRoute := r.Group("/events").Use(authMiddleware)
	{
		eventRoute.GET("/", eventHandler.GetEvents)
		eventRoute.POST("/", eventHandler.CreateEvent)
		eventRoute.GET("/:id", eventHandler.GetEvent)
		eventRoute.PUT("/:id", eventHandler.UpdateEvent)
		eventRoute.DELETE("/:id", eventHandler.DeleteEvent)
		eventRoute.DELETE("/", eventHandler.DeleteMultipleEvent)
		eventRoute.GET("/get-created-events", eventHandler.GetCreatedEvent)
		eventRoute.PATCH("/restore", eventHandler.RestoreEvents)
		eventRoute.GET("/get-deleted-events", eventHandler.GetTrashedEvent)
		eventRoute.PATCH("/favourite/:id", eventHandler.FavouriteEvent)
		eventRoute.PATCH("/unfavourite/:id", eventHandler.UnFavouriteEvent)
		eventRoute.GET("/get-favourite-events", eventHandler.GetFavouriteEvent)
		eventRoute.PATCH("/make-events-private", eventHandler.MakeEventPrivate)
		eventRoute.PATCH("/make-events-public", eventHandler.MakeEventPublic)
		eventRoute.PATCH("/apply-coupons/:id", eventHandler.ApplyCoupons)
		eventRoute.PATCH("/remove-coupons/:id", eventHandler.RemoveCoupons)
	}
}
