package http

import (
	"errors"
	"github.com/QuocAnh189/GoBin/logger"
	"gohub/domains/users/dto"
	"gohub/domains/users/service"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"gohub/pkg/utils"
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

//		@Summary	 Retrieve a list of users
//	 @Description Fetches a paginated list of users based on the provided filter parameters.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of users"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users [get]
func (u *UserHandler) GetUsers(c *gin.Context) {
	var req dto.ListUserReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")

	var res dto.ListUserRes
	users, pagination, err := u.service.GetUsers(c, &req, userId)
	if err != nil {
		logger.Error("Failed to get list products: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res.User, &users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve user profile
//	 @Description Fetches the details of the currently authenticated user.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User profile retrieved successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/profile [get]
func (u *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("userId")

	if userID == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	user, calculation, err := u.service.GetProfile(c, userID)
	if err != nil {
		logger.Error(err.Error())
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.User
	utils.MapStruct(&res, &user)
	res.TotalEvent = calculation.TotalEvent
	res.TotalFollower = calculation.TotalFollower
	res.TotalFollowing = calculation.TotalFollowing
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a user by its ID
//	 @Description Successfully retrieved the user
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the user"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{userId} [get]
func (u *UserHandler) GetUserById(c *gin.Context) {
	userId := c.Param("id")

	user, calculation, err := u.service.GetUserById(c, userId)
	if err != nil {
		logger.Error("Failed to get user: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.User
	utils.MapStruct(&res, user)
	res.TotalEvent = calculation.TotalEvent
	res.TotalFollower = calculation.TotalFollower
	res.TotalFollowing = calculation.TotalFollowing
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new user
//	 @Description Creates a new user based on the provided details. The request must include multipart form data.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"User created successfully"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	user, err := u.service.CreateUser(c, &req)
	if err != nil {
		logger.Error("Failed to create user ", err.Error())
		switch err.Error() {
		case messages.EmailAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.EmailAlreadyExists)
		case messages.UserNameAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.UserNameAlreadyExists)
		case messages.PhoneNumberAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.PhoneNumberAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create user")
		}
		return
	}

	var res dto.User
	utils.MapStruct(&res, &user)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing user
//	 @Description Updates the details of an existing user based on the provided user ID and update information.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User updated successfully"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{userId} [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var req dto.UpdateUserReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.Param("id")
	user, err := u.service.UpdateUser(c, userId, &req)
	if err != nil {
		logger.Error("Failed to update user ", err.Error())
		switch err.Error() {
		case messages.EmailAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.EmailAlreadyExists)
		case messages.UserNameAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.UserNameAlreadyExists)
		case messages.PhoneNumberAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.PhoneNumberAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update user")
		}
		return
	}

	var res dto.User
	utils.MapStruct(&res, &user)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Change a user's password
//	 @Description Changes the password of an existing user based on the provided user ID and new password information.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User updated successfully"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{userId}/change-password [patch]
func (u *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePassword
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")

	err := u.service.ChangePassword(c, userId, &req)
	if err != nil {
		switch err.Error() {
		case messages.WrongPassword:
			response.Error(c, http.StatusConflict, err, messages.AccountOrPasswordWrong)
		case messages.UserNotFound:
			response.Error(c, http.StatusNotFound, err, messages.UserNotFound)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	response.JSON(c, http.StatusOK, "Change password successfully")
}

//		@Summary	 Retrieve a list of a user by its ID
//	 @Description Fetches a paginated list of followers based on the provided user ID and filter parameters.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of followers"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{userId}/followers [get]
func (u *UserHandler) GetFollowers(c *gin.Context) {
	var req dto.ListUserReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.Param("id")
	_, _, err := u.service.GetUserById(c, userId)
	if err != nil {
		logger.Error("Failed to get user detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListUserRes
	users, pagination, err := u.service.GetFollowers(c, &req, userId)
	if err != nil {
		logger.Error("Failed to get list followers: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res.User, &users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a list of a user by its ID
//	 @Description Fetches a paginated list of following users based on the provided user ID and filter parameters.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of following users"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{userId}/followings [get]
func (u *UserHandler) GetFollowing(c *gin.Context) {
	var req dto.ListUserReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.Param("id")
	_, _, err := u.service.GetUserById(c, userId)
	if err != nil {
		logger.Error("Failed to get user detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListUserRes
	users, pagination, err := u.service.GetFollowing(c, &req, userId)
	if err != nil {
		logger.Error("Failed to get list followings: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res.User, &users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Follow a user
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User followed successfully"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/follow/{followedUserId} [patch]
func (u *UserHandler) FollowUser(c *gin.Context) {
	var req dto.FollowerUserReq
	req.FollowerId = c.GetString("userId")
	req.FolloweeId = c.Param("followeeId")

	_, _, err := u.service.GetUserById(c, req.FollowerId)
	if err != nil {
		logger.Error("Failed to get follower: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	_, _, err = u.service.GetUserById(c, req.FolloweeId)
	if err != nil {
		logger.Error("Failed to get followee: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	isFollower := u.service.IsFollower(c, &req)
	if isFollower {
		logger.Error("Failed to follow user: ", err)
		response.Error(c, http.StatusUnprocessableEntity, err, "This user is already followed by you")
		return
	}

	err = u.service.FollowUser(c, &req)
	if err != nil {
		logger.Error("Failed to follow user: ", err)
	}
	response.JSON(c, http.StatusOK, "Follow user successfully")
}

//		@Summary	 UnFollow a user
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User followed successfully"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/unfollow/{followedUserId} [patch]
func (u *UserHandler) UnfollowUser(c *gin.Context) {
	var req dto.FollowerUserReq
	req.FollowerId = c.GetString("userId")
	req.FolloweeId = c.Param("followeeId")

	_, _, err := u.service.GetUserById(c, req.FollowerId)
	if err != nil {
		logger.Error("Failed to get follower: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	_, _, err = u.service.GetUserById(c, req.FolloweeId)
	if err != nil {
		logger.Error("Failed to get followee: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	isFollower := u.service.IsFollower(c, &req)
	if !isFollower {
		logger.Error("Failed to follow user: ", err)
		response.Error(c, http.StatusBadRequest, err, "You have never followed this user")
		return
	}

	err = u.service.UnfollowUser(c, &req)
	if err != nil {
		logger.Error("Failed to unfollow user: ", err)
	}
	response.JSON(c, http.StatusOK, "UnFollow user successfully")
}

//		@Summary	 Check Follower
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/check-follower/{followedUserId} [get]
func (u *UserHandler) CheckFollower(c *gin.Context) {
	var req dto.FollowerUserReq
	req.FollowerId = c.GetString("userId")
	req.FolloweeId = c.Param("followeeId")

	result, err := u.service.CheckFollower(c, &req)
	if err != nil {
		logger.Error("Failed to check follower: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
		return
	}

	response.JSON(c, http.StatusOK, result)
}

//		@Summary	 Get Invitations
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/invitation [get]
func (u *UserHandler) GetInvitations(c *gin.Context) {
	var req dto.ListInvitationReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	inviteeId := c.GetString("userId")

	var res dto.ListInvitationRes
	invitations, pagination, err := u.service.GetInvitations(c, &req, inviteeId)
	if err != nil {
		logger.Error("Failed to get invitations: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
		return
	}

	utils.MapStruct(&res.Invitations, &invitations)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Invitate User
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/invitation [patch]
func (u *UserHandler) InviteUsers(c *gin.Context) {
	var req dto.InviteUsers
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	if err := u.service.InviteUsers(c, &req, userId); err != nil {
		logger.Error("Failed to invite users: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
		return
	}

	response.JSON(c, http.StatusOK, "Invite Successfully")
}

//		@Summary	 Check Follower
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/check-invitation/{inviteeId} [get]
func (u *UserHandler) CheckInvitation(c *gin.Context) {
	var req dto.CheckInvitationReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	//inviteeId := c.Param("inviteeId")
	userId := c.GetString("userId")

	result, err := u.service.CheckInvitation(c, &req, userId)
	if err != nil {
		logger.Error("Failed to check follower: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
		return
	}

	response.JSON(c, http.StatusOK, result)
}

//		@Summary	 Get Notification Following
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/notification-following [get]
func (u *UserHandler) GetNotificationFollowings(c *gin.Context) {
	var req dto.ListNotificationReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	inviteeId := c.GetString("userId")

	var res dto.ListNotificationFollowingRes
	results, pagination, err := u.service.GetNotificationFollowings(c, &req, inviteeId)
	if err != nil {
		logger.Error("Failed to get invitations: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
		return
	}

	utils.MapStruct(&res.Notifications, &results)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Get Notification Order
//	 @Description Allows the authenticated user to follow another user by specifying the followed user's ID.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"true or false"
//		@Failure	 400	{object}	response.Response	"BadRequest - Invalid input or request data"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/notification-order [get]
func (u *UserHandler) GetNotificationOrders(c *gin.Context) {
}
