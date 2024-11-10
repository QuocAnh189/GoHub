package http

import (
	"gohub/domains/roles/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)
type RoleHandler struct {
	service service.IRoleService
}

func NewRoleHandler(service service.IRoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

//	@Summary	 Add a function to a role
//  @Description Adds a specific function to the role identified by the role ID.
//	@Tags		 Roles
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of users"
//	@Failure	 400	{object}	response.Response	"Bad Request - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Role or function with the specified IDs not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/roles/{roleId}/add-function/{functionId} [post]
func (h *RoleHandler) AddFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test AddFunction route"})
}

//	@Summary	 Remove a function from a role
//  @Description Removes a specific function from the role identified by the role ID.
//	@Tags		 Roles
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of users"
//	@Failure	 400	{object}	response.Response	"Bad Request - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Role or function with the specified IDs not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/roles/{roleId}/remove-function/{functionId} [post]
func (h *RoleHandler) RemoveFunction(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test RemoveFunction route"})
}