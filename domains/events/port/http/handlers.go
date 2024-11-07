package http

import (
	"gohub/domains/events/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service service.IEventService
}

func NewEventHandler(service service.IEventService) *EventHandler {
	return &EventHandler{service: service}
}

//	@Summary	 Retrieve a list of events
//  @Description Fetches a paginated list of events based on the provided filter parameters.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of events"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events [get]
func (h *EventHandler) GetEvents(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetEvents route"})
}

//	@Summary	 Create a new event
//  @Description Creates a new event with the provided details.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event created successfully"
//	@Success	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateEvent route"})
}

//	@Summary	 Retrieve a event by its ID
//  @Description Fetches the details of a specific event based on the provided event ID.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the event"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/{eventId} [get]
func (h *EventHandler) GetEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetEvent route"})
}

//	@Summary	 Update an existing event
//  @Description Updates the details of an existing event based on the provided event ID and update information.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the event"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/{eventId} [put]
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateEvent route"})
}

//	@Summary	 Delete an existing event
//  @Description Deletes an existing event based on the provided event ID.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event deleted successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/{eventId} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test DeleteEvent route"})
}

//	@Summary	 Retrieve created events
//  @Description Fetches a paginated list of events created by the user, based on the provided pagination filter.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved events"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/get-created-events [get]
func (h *EventHandler) GetCreatedEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetCreatedEvent route"})
}	

//	@Summary	 Permanently delete an event
//  @Description Permanently deletes an existing event based on the provided event ID.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event permanently deleted successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/delete-permanently/{eventId} [delete]
func (h *EventHandler) DeletePermanentlyEvent(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test DeletePermanentlyEvent route"})
}

//	@Summary	 Restore deleted events
//  @Description Restores a list of deleted events based on the provided event IDs.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event permanently deleted successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/restore [patch]
func (h *EventHandler) RestoreEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test RestoreEvent route"})
}

//	@Summary	 Retrieve deleted events
//  @Description Fetches a paginated list of events that have been deleted, based on the provided pagination filter.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event permanently deleted successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/get-delete-events [get]
func (h *EventHandler) GetTrashedEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetTrashedEvent route"})
}	

//	@Summary	 Mark anb event as favorite
//  @Description Marks an existing event as a favourite based on the provided event ID.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//	@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/favourite/{eventId} [patch]
func (h *EventHandler) FavouriteEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test FavouriteEvent route"})
}

//	@Summary	 UnMark anb event as favorite
//  @Description Removes an event from the user's favourites based on the provided event ID.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//	@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/unfavourite/{eventId} [patch]
func (h *EventHandler) UnfavouriteEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UnfavouriteEvent route"})
}

//	@Summary	 Retrieve favourite events
//  @Description Fetches a paginated list of events marked as favourites by the user, based on the provided pagination filter.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Event marked as favourite successfully"
//	@Success	 400	{object}	response.Response	"BadRequest - Invalid request data"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Success	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/get-favourite-events [get]
func (h *EventHandler) GetFavouriteEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFavouriteEvent route"})
}

//	@Summary	 Make events private
//  @Description Changes the visibility of specified events to private based on the provided event IDs.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Events marked as private successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/make-events-private [patch]
func (h *EventHandler) MakeEventPrivate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test MakeEventPrivate route"})
}

//	@Summary	 Make events public
//  @Description Changes the visibility of specified events to public based on the provided event IDs.
//	@Tags		 Events
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Events marked as public successfully"
//	@Success	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Success	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/events/make-events-public [patch]
func (h *EventHandler) MakeEventPublic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test MakeEventPublic route"})
}

