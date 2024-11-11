package http

import (
	"gohub/domains/permissions/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	service service.IPermissionService
}

func NewPermissionHandler(service service.IPermissionService) *PermissionHandler {
	return &PermissionHandler{service: service}
}

//	@Summary	 Retrieve all permissions for a function
//  @Description Fetches a list of all permissions associated with the specified function ID.
//	@Tags		 Permissions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of permissions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/permissions [get]
func (p *PermissionHandler) GetPermissions(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetPermission route"})
}

//	@Summary	 Retrieve permissions categorized by roles
//  @Description Fetches a list of permissions categorized by roles.
//	@Tags		 Permissions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of permissions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/permissions/roles [get]
func (p *PermissionHandler) GetPermissionsByRoles(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetPermissionByRoles route"})
}

//	@Summary	 Retrieve permissions by user
//  @Description Fetches a list of permissions by user id.
//	@Tags		 Permissions
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of permissions"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/permissions/get-by-user/{userId} [get]
func (p *PermissionHandler) GetPermissionsByUsers(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetPermissionByUser route"})
}