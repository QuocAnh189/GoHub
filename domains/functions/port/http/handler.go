package http

import (
	"gohub/domains/functions/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FunctionHandler struct {
	service service.IFunctionService
}

func NewFunctionHandler(service service.IFunctionService) *FunctionHandler {
	return &FunctionHandler{service: service}
}

//	@Summary	 Create a new function
//  @Description Creates a new function based on the provided details.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 201	{object}	response.Response	"Function created successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions [post]
func (h *FunctionHandler) CreateFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateFunction route"})
}

//	@Summary	 Retrieve a list of functions
//  @Description Fetches a list of all available functions.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions [get]
func (h *FunctionHandler) GetFunctions(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFunctions route"})
}

//	@Summary	 Retrieve a function by its ID
//  @Description Fetches the details of a specific function based on the provided function ID.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Function with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions/{functionId} [get]
func (h *FunctionHandler) GetFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFunction route"})
}

//	@Summary	 Update an existing function
//  @Description Updates the details of an existing function based on the provided function ID and update information.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Function with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions/{functionId} [put]
func (h *FunctionHandler) UpdateFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateFunction route"})
}

//	@Summary	 Delete a function
//  @Description Deletes the function with the specified ID.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Function with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions/{functionId} [delete]
func (h *FunctionHandler) DeleteFunction(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test DeleteFunction route"})
}

//	@Summary	 Enable a command for a function
//  @Description Enables a specific command for the function identified by the function ID.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 400	{object}	response.Response	"Bad Request - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Function with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions/{functionId}/enable-command/{commandId} [post]
func (h *FunctionHandler) EnableCommand(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test EnableCommand route"})
}

//	@Summary	 Disable command for a function
//  @Description Disables a specific command for the function identified by the function ID.
//	@Tags		 Functions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of functions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Function not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - Function does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Function with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/functions/{functionId}/disable-command/{commandId} [post]
func (h *FunctionHandler) DisableCommand(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test EnableCommand route"})
}

