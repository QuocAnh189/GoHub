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

func (u *UserHandler) GetUsers(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetUsers route"})
}

func (u *UserHandler) GetUser(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test GetUser route"})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateUser route"})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateUser route"})
}

func (u *UserHandler) ChangePassword(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ChangePassword route"})
}

func (u *UserHandler) GetFollowers(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFollowers route"})
}

func (u *UserHandler) GetFollowing(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetFollowing route"})
}

func (u *UserHandler) FollowUser(c *gin.Context) {	
	response.JSON(c, http.StatusOK, gin.H{"message": "Test FollowUser route"})
}

func (u *UserHandler) UnfollowUser(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UnfollowUser route"})
}