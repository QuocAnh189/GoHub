package http

import (
	"gohub/domains/commands/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommandHandler struct {
	service service.ICommandService
}

func NewCommandHandler(service service.ICommandService) *CommandHandler {
    return &CommandHandler{service: service}
}

//	@Summary	 Retrieve commands associated with specific function
//  @Description Fetches a list of commands that are associated with the specified function ID.
//	@Tags		 Commands
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of commands"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/commands/get-in-function/{functionId} [get]
func (h *CommandHandler) GetInFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetInFunction route"})
}

//	@Summary	 Retrieve commands not associated with a specific function
//  @Description Fetches a list of commands that are not associated with the specified function ID.
//	@Tags		 Commands
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of commands"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/commands/get-not-in-function/{functionId} [get]
func (h *CommandHandler) GetNotInFunction(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test GetNotInFunction route"})
}