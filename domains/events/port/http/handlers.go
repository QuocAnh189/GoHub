package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"gohub/domains/events/dto"
	"gohub/domains/events/service"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service service.IEventService
}

func NewEventHandler(service service.IEventService) *EventHandler {
	return &EventHandler{service: service}
}

//		@Summary	 Retrieve a list of events
//	 @Description Fetches a paginated list of events based on the provided filter parameters.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of events"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events [get]
func (h *EventHandler) GetEvents(c *gin.Context) {
	var req dto.ListEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	events, pagination, err := h.service.GetEvents(c, &req)
	if err != nil {
		logger.Error("Failed to get events: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get events")
		return
	}

	var res dto.ListEventRes
	utils.MapStruct(&res.Events, &events)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new event
//	 @Description Creates a new event with the provided details.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event created successfully"
//		@Success	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req dto.CreateEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	event, err := h.service.CreateEvent(c, &req)
	if err != nil {
		logger.Error("Failed to create event ", err.Error())
		switch err.Error() {
		case messages.EventNameAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.EventNameAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create event")
		}
		return
	}

	var res dto.Event
	utils.MapStruct(&res, &event)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a event by its ID
//	 @Description Fetches the details of a specific event based on the provided event ID.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the event"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/{eventId} [get]
func (h *EventHandler) GetEvent(c *gin.Context) {
	eventId := c.Param("id")
	event, err := h.service.GetEventById(c, eventId)
	if err != nil {
		logger.Error("Failed to get event: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
	}

	var res dto.Event
	utils.MapStruct(&res, &event)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing event
//	 @Description Updates the details of an existing event based on the provided event ID and update information.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the event"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/{eventId} [put]
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	eventId := c.Param("id")
	var req dto.UpdateEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	req.UserId = c.GetString("userId")

	event, err := h.service.UpdateEvent(c, eventId, &req)
	if err != nil {
		logger.Error("Failed to update event ", err.Error())
		switch err.Error() {
		case messages.CategoryNameExists:
			response.Error(c, http.StatusConflict, err, messages.CategoryNameExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update event")
		}
		return
	}

	var res dto.Event
	utils.MapStruct(&res, &event)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete an existing event
//	 @Description Deletes an existing event based on the provided event ID.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event deleted successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/{eventId} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	eventId := c.Param("id")

	err := h.service.DeleteEvent(c, eventId)

	if err != nil {
		logger.Error("Failed to delete event: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete event successfully")
}

//		@Summary	 Delete an existing multiple event
//	 @Description Deletes an existing event based on the provided event ID.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Events deleted successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/{eventId} [delete]
func (h *EventHandler) DeleteMultipleEvent(c *gin.Context) {
	var req dto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.DeleteEvents(c, &req)

	if err != nil {
		logger.Error("Failed to delete events: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete events successfully")
}

//		@Summary	 Retrieve created events
//	 @Description Fetches a paginated list of events created by the user, based on the provided pagination filter.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved events"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/get-created-events [get]
func (h *EventHandler) GetCreatedEvent(c *gin.Context) {
	var req dto.ListEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	events, pagination, err := h.service.GetCreatedEvent(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get events: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get events")
		return
	}

	var res dto.ListEventRes
	utils.MapStruct(&res.Events, &events)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Restore deleted events
//	 @Description Restores a list of deleted events based on the provided event IDs.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event permanently deleted successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/restore [patch]
func (h *EventHandler) RestoreEvents(c *gin.Context) {
	var req dto.RestoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := h.service.RestoreEvents(c, &req)

	if err != nil {
		logger.Error("Failed to events category: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Restore events successfully")
}

//		@Summary	 Retrieve deleted events
//	 @Description Fetches a paginated list of events that have been deleted, based on the provided pagination filter.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event permanently deleted successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/get-delete-events [get]
func (h *EventHandler) GetTrashedEvent(c *gin.Context) {
	var req dto.ListEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	events, pagination, err := h.service.GetTrashedEvent(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get events: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get events")
		return
	}

	var res dto.ListEventRes
	utils.MapStruct(&res.Events, &events)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Mark anb event as favorite
//	 @Description Marks an existing event as a favourite based on the provided event ID.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//		@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/favourite/{eventId} [patch]
func (h *EventHandler) FavouriteEvent(c *gin.Context) {
	req := dto.CreateEventFavouriteReq{
		UserID:  c.GetString("userId"),
		EventId: c.Param("id"),
	}

	err := h.service.FavouriteEvent(c, &req)
	if err != nil {
		logger.Error("Failed to create event ", err.Error())
		switch err.Error() {
		case messages.EventFavouriteAlreadyExists:
			response.Error(c, http.StatusInternalServerError, err, messages.EventFavouriteAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create category")
		}
		return
	}

	response.JSON(c, http.StatusOK, "Favourite event successfully")
}

//		@Summary	 UnMark anb event as favorite
//	 @Description Removes an event from the user's favourites based on the provided event ID.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//		@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/unfavourite/{eventId} [patch]
func (h *EventHandler) UnFavouriteEvent(c *gin.Context) {
	req := dto.CreateEventFavouriteReq{
		UserID:  c.GetString("userId"),
		EventId: c.Param("id"),
	}

	err := h.service.UnFavouriteEvent(c, &req)
	if err != nil {
		logger.Error("Failed to create event ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create category")
		}
		return
	}

	response.JSON(c, http.StatusOK, "UnFavourite event successfully")
}

//		@Summary	 Retrieve favourite events
//	 @Description Fetches a paginated list of events marked as favourites by the user, based on the provided pagination filter.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//		@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/get-favourite-events [get]
func (h *EventHandler) GetFavouriteEvent(c *gin.Context) {
	var req dto.ListEventReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	events, pagination, err := h.service.GetFavouriteEvent(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get events: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get events")
		return
	}

	var res dto.ListEventFavouriteRes
	utils.MapStruct(&res.Events, &events)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Make events private
//	 @Description Changes the visibility of specified events to private based on the provided event IDs.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Events marked as private successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/make-events-private [patch]
func (h *EventHandler) MakeEventPrivate(c *gin.Context) {
	var req dto.MakeEventPublicOrPrivateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	req.UserId = c.GetString("userId")

	err := h.service.MakeEventPrivate(c, &req)
	if err != nil {
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
	}

	response.JSON(c, http.StatusOK, "Make events private successfully")
}

//		@Summary	 Make events public
//	 @Description Changes the visibility of specified events to public based on the provided event IDs.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Events marked as public successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/make-events-public [patch]
func (h *EventHandler) MakeEventPublic(c *gin.Context) {
	var req dto.MakeEventPublicOrPrivateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	req.UserId = c.GetString("userId")

	err := h.service.MakeEventPublic(c, &req)
	if err != nil {
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
	}

	response.JSON(c, http.StatusOK, "Make events public successfully")
}

//		@Summary	 Apply Coupons for event
//	 @Description Changes the visibility of specified events to public based on the provided event IDs.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Events marked as public successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/apply-coupons/{eventId} [patch]
func (h *EventHandler) ApplyCoupons(c *gin.Context) {
	var req dto.ApplyCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
	}

	eventId := c.Param("id")
	err := h.service.ApplyCoupons(c, eventId, &req)
	if err != nil {
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
			return
		}
	}

	response.JSON(c, http.StatusOK, "Apply coupons successfully")
}

//		@Summary	 Remove Coupons from event
//	 @Description Changes the visibility of specified events to public based on the provided event IDs.
//		@Tags		 Events
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Events marked as public successfully"
//		@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/events/remove-coupons/{eventId} [patch]
func (h *EventHandler) RemoveCoupons(c *gin.Context) {
	var req dto.RemoveCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
	}

	eventId := c.Param("id")
	err := h.service.RemoveCoupons(c, eventId, &req)
	if err != nil {
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
			return
		}
	}

	response.JSON(c, http.StatusOK, "Remove coupons successfully")
}
