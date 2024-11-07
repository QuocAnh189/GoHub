package http

import (
	"gohub/domains/users/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

//	@Summary	 Retrieve a list of users
//  @Description Fetches a paginated list of users based on the provided filter parameters.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of users"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users [get]
func (u *UserHandler) GetUsers(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetUsers route"})
}

//	@Summary	 Retrieve a user by its ID
//  @Description Successfully retrieved the user
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the user"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/{userId} [get]
func (u *UserHandler) GetUser(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test GetUser route"})
}

//	@Summary	 Create a new user
//  @Description Creates a new user based on the provided details. The request must include multipart form data.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 201	{object}	response.Response	"User created successfully"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateUser route"})
}

//	@Summary	 Update an existing user
//  @Description Updates the details of an existing user based on the provided user ID and update information.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"User updated successfully"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/{userId} [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateUser route"})
}

//	@Summary	 Change a user's password
//  @Description Changes the password of an existing user based on the provided user ID and new password information.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"User updated successfully"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/{userId}/change-password [patch]
func (u *UserHandler) ChangePassword(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ChangePassword route"})
}

//	@Summary	 Retrieve a list of a user by its ID
//  @Description Fetches a paginated list of followers based on the provided user ID and filter parameters.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of followers"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/{userId}/followers [get]
func (u *UserHandler) GetFollowers(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFollowers route"})
}

//	@Summary	 Retrieve a list of a user by its ID
//  @Description Fetches a paginated list of following users based on the provided user ID and filter parameters.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the list of following users"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/{userId}/following-users [get]
func (u *UserHandler) GetFollowing(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFollowing route"})
}

//	@Summary	 Follow a user
//  @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"User followed successfully"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/follow/{followedUserId} [patch]
func (u *UserHandler) FollowUser(c *gin.Context) {	
	response.JSON(c, http.StatusOK, gin.H{"message": "Test FollowUser route"})
}

//	@Summary	 UnFollow a user
//  @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//	@Tags		 Users
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"User followed successfully"
//	@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/users/unfollow/{followedUserId} [patch]
func (u *UserHandler) UnfollowUser(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UnfollowUser route"})
}