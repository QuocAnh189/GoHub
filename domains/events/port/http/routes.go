package http

import (
	"gohub/database"
	"gohub/domains/events/repository"
	"gohub/domains/events/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	eventRepository := repository.NewEventRepository(sqlDB)
	eventService := service.NewEventService(validator, eventRepository)
	eventHandler := NewEventHandler(eventService)

	eventRoute := r.Group("/events")
	{
		eventRoute.GET("/", eventHandler.GetEvents)
		eventRoute.POST("/", eventHandler.CreateEvent)
		eventRoute.GET("/:id", eventHandler.GetEvent)
		eventRoute.PUT("/:id", eventHandler.UpdateEvent)
		eventRoute.DELETE("/:id", eventHandler.DeleteEvent)
		eventRoute.GET("/get-created-events", eventHandler.GetCreatedEvent)
		eventRoute.DELETE("/delete-permanently/:id", eventHandler.DeletePermanentlyEvent)
		eventRoute.PATCH("/restore", eventHandler.RestoreEvent)
		eventRoute.GET("/get-deleted-events", eventHandler.GetTrashedEvent)
		eventRoute.PATCH("/favourite/:id", eventHandler.FavouriteEvent)
		eventRoute.PATCH("/unfavourite/:id", eventHandler.UnfavouriteEvent)
		eventRoute.GET("/get-favourite-events", eventHandler.GetFavouriteEvent)
		eventRoute.PATCH("/make-events-private/:id", eventHandler.MakeEventPrivate)
		eventRoute.PATCH("/make-events-public/:id", eventHandler.MakeEventPublic)
	}
}