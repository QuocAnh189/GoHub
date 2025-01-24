package http

import (
	"gohub/domains/conversations/dto"
	"gohub/domains/conversations/service"
	"gohub/internal/libs/logger"
	"gohub/pkg/response"
	"gohub/pkg/utils"
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

//		@Summary	 Retrieves a list of conversations by the event
//	 @Description Fetches a paginated list of conversations created by the event, based on the provided pagination filter.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of conversations"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/get-by-organizer/{organizerId} [get]
func (h *ConversationHandler) GetConversationsByOrganizer(c *gin.Context) {
	var req *dto.ListConversationReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	organizerId := c.Param("organizerId")
	conversations, pagination, err := h.service.GetConversationsOrganizer(c, organizerId, req)
	if err != nil {
		logger.Error("Failed to get conversations: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListConversationByOrganizerRes
	utils.MapStruct(&res.Conversation, &conversations)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieves a list of conversations by the user
//	 @Description Fetches a paginated list of conversations created by the user, based on the provided pagination filter.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of conversations"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/get-by-user/{userId} [get]
func (h *ConversationHandler) GetConversationsByUser(c *gin.Context) {
	var req *dto.ListConversationReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.Param("userId")
	conversations, pagination, err := h.service.GetConversationsByUser(c, userId, req)
	if err != nil {
		logger.Error("Failed to get conversations: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListConversationByUserRes
	utils.MapStruct(&res.Conversation, &conversations)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieves a list of messages by the conversation
//	 @Description Fetches a paginated list of messages created by the conversation, based on the provided pagination filter.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of messages"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/{conversationId}/messages [get]
func (h *ConversationHandler) GetMessagesByConversation(c *gin.Context) {
	var req *dto.ListMessageReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	conversationId := c.Param("id")
	messages, pagination, err := h.service.GetMessagesByConversation(c, conversationId, req)
	if err != nil {
		logger.Error("Failed to get messages: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListMessageRes
	utils.MapStruct(&res.Messages, messages)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create message in conversation
//	 @Description Fetches a paginated list of messages created by the conversation, based on the provided pagination filter.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of messages"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/{conversationId}/messages [post]
func (h *ConversationHandler) CreateMessage(c *gin.Context) {
	var req dto.CreateMessageReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	message, err := h.service.CreateMessage(c, &req)
	if err != nil {
		logger.Error("Failed to create message ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create message")
		}
		return
	}

	var res dto.Message
	utils.MapStruct(&res, &message)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing message
//	 @Description Updates the details of an existing category based on the provided category ID and update information.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/{conversationId}/messages/{messageId} [put]
func (h *ConversationHandler) UpdateMessage(c *gin.Context) {
	messageId := c.Param("messageId")
	var req dto.UpdateMessageReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	message, err := h.service.UpdateMessage(c, messageId, &req)
	if err != nil {
		logger.Error("Failed to update message", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update category")
		}
		return
	}

	var res dto.Message
	utils.MapStruct(&res, &message)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete a message
//	 @Description Deletes the category with the specified ID.
//		@Tags		 Conversations
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/conversations/{conversationId}/messages/{messageId} [delete]
func (h *ConversationHandler) DeleteMessage(c *gin.Context) {
	messageId := c.Param("messageId")

	message, err := h.service.DeleteMessage(c, messageId)

	if err != nil {
		logger.Error("Failed to delete message: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.Message
	utils.MapStruct(&res, &message)
	response.JSON(c, http.StatusOK, res)
}
