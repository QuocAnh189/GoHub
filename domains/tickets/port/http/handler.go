package http

import (
	"github.com/gin-gonic/gin"
	"gohub/domains/tickets/dto"
	"gohub/domains/tickets/service"
	"gohub/internal/libs/logger"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"
)

type TicketHandler struct {
	service service.ITicketService
}

func NewTicketHandler(service service.ITicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

//		@Summary	 Retrieve an tickets by created
//	 @Description Fetches the details of a specific expense based on the provided expense ID.
//		@Tags		 Tickets
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/tickets/get-created-tickets [get]
func (h *TicketHandler) GetTicketByCreated(c *gin.Context) {
	var req dto.ListTicketReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	tickets, pagination, err := h.service.GetCreatedTickets(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get list tickets: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListTicketRes
	utils.MapStruct(&res.Ticket, &tickets)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}
