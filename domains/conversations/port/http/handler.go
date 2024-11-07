package http

import (
	"gohub/domains/conversations/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	service service.IConversationService
}

func NewConversationHandler(service service.IConversationService) *ConversationHandler {
	return &ConversationHandler{
		service: service,
	}
}

//	@Summary	 Retrieves a list of conversations by the event
//  @Description Fetches a paginated list of conversations created by the event, based on the provided pagination filter.
//	@Tags		 Conversations
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of conversations"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/conversations/get-by-event/{eventId} [get]
func (h *ConversationHandler) GetConversationsByEvent(c *gin.Context)  {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ConversationsByEvent route"})
}

//	@Summary	 Retrieves a list of conversations by the user
//  @Description Fetches a paginated list of conversations created by the user, based on the provided pagination filter.
//	@Tags		 Conversations
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of conversations"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/conversations/get-by-user/{userId} [get]
func (h *ConversationHandler) GetConversationsByUser(c *gin.Context)  {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ConversationsByUser route"})
}

//	@Summary	 Retrieves a list of messages by the conversation
//  @Description Fetches a paginated list of messages created by the conversation, based on the provided pagination filter.
//	@Tags		 Conversations
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of messages"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/conversations/{conversationId}/messages [get]
func (h *ConversationHandler) GetMessagesByConversation(c *gin.Context)  {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test MessagesByConversation route"})
}

