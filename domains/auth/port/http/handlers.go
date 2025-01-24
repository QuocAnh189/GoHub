package http

import (
	"errors"
	"fmt"
	"gohub/domains/auth/dto"
	"gohub/domains/auth/service"
	"gohub/internal/libs/logger"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.IAuthService
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

//		@Summary	 Validate user credentials
//	 @Description Validate user information session one when signup
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.ValidateUserReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"User credentials are valid"
//		@Failure	 400	{object}	response.Response	"Invalid user credentials or request data"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/validate-user [post]
func (auth *AuthHandler) ValidateUser(c *gin.Context) {
	var req dto.ValidateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	err := auth.service.ValidateUser(c, &req)
	if err != nil {
		logger.Error(err.Error())
		switch err.Error() {
		case messages.EmailAlreadyExists:
			response.Error(c, http.StatusUnprocessableEntity, err, messages.EmailAlreadyExists)
		case messages.UserNameAlreadyExists:
			response.Error(c, http.StatusUnprocessableEntity, err, messages.UserNameAlreadyExists)
		case messages.PhoneNumberAlreadyExists:
			response.Error(c, http.StatusUnprocessableEntity, err, messages.PhoneNumberAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	response.JSON(c, http.StatusOK, "Validation successfully")
}

//		@Summary	 Signup a new user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.SignUpReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"User successfully registered"
//		@Failure	 400	{object}	response.Response	"Invalid user input"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signup [post]
func (auth *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, err := auth.service.SignUp(c, &req)
	if err != nil {
		logger.Error(err.Error())
		switch err.Error() {
		case messages.EmailAlreadyExists:
			response.Error(c, http.StatusUnprocessableEntity, err, messages.EmailAlreadyExists)
		case messages.UserNameAlreadyExists:
			response.Error(c, http.StatusUnprocessableEntity, err, messages.UserNameAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	var res dto.SignUpRes
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Signin a user
//	 @Description Authenticates the user based on the provided credentials and returns a sign-in response if successful.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.SignInReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"Successfully signed in"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid credentials"
//		@Failure	 404	{object}	response.Response	"Not Found - User not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signin [post]
func (auth *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, err := auth.service.SignIn(c, &req)
	if err != nil {
		logger.Error("Failed to login ", err.Error())
		switch err.Error() {
		case messages.AccountOrPasswordWrong:
			response.Error(c, http.StatusConflict, err, messages.AccountOrPasswordWrong)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	var res dto.SignInRes
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Signout a new user
//	 @Description Signs out the current user, invalidating their session or authentication token.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully signed out"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signout [post]
func (auth *AuthHandler) SignOut(c *gin.Context) {
	jwtToken := c.GetString("token")
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	err := auth.service.SignOut(c, cleanJWT)

	if err != nil {
		logger.Error("Failed to logout", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to logout")
		return
	}

	res := dto.SignOutRes{
		Message: "Logout successful",
	}
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Signin a new user with external credentials
//	 @Description Authenticates the user using an external authentication provider (e.g., Google, Facebook) and returns a login response if successful.
//		@Tags		 Auth
//		@Produce	 json
//		@Param 	     provider query string true "Authentication provider"
//		@Param       returnUrl query string true "Redirect URL after authentication"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/external-login [post]
func (auth *AuthHandler) ExternalSignIn(c *gin.Context) {
	auth.service.ExternalSignIn(c.Writer, c.Request)
}

//		@Summary	 Callback endpoint for external authentication
//	 @Description Handles the callback from an external authentication provider and processes the authentication result.
//		@Tags		 Auth
//		@Produce	 json
//		@Param       returnUrl query string true "Redirect URL after authentication"
//		@Failure	 400	{object}	response.Response	"Invalid user credentials or request data"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/external-auth-callback [get]
func (auth *AuthHandler) ExternalCallback(c *gin.Context) {
	user, err := auth.service.ExternalCallback(c.Writer, c.Request)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err, err.Error())
		return
	}
	fmt.Println(user)
	http.Redirect(c.Writer, c.Request, "http://localhost:3000/organization", http.StatusFound)
	//response.JSON(c, http.StatusOK, user)
}

//		@Summary	 Refresh user authentication token
//	 @Description Refreshes the user's authentication token based on the provided refresh token credentials.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.RefreshTokenReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"Successfully refreshed the token"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or expired refresh token"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/refresh-token [post]
func (auth *AuthHandler) RefreshToken(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		response.Error(c, http.StatusUnauthorized, errors.New("unauthorized"), "Unauthorized")
		return
	}

	accessToken, err := auth.service.RefreshToken(c, userId)
	if err != nil {
		logger.Error("Failed to refresh token", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	res := dto.RefreshTokenRes{
		AccessToken: accessToken,
	}
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Initiate password recovery
//	 @Description Sends a password recovery email to the user based on the provided email address.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.ForgotPasswordReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"Password recovery email sent successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or missing credentials"
//		@Failure	 404	{object}	response.Response	"Not Found - User with the provided email address not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/forgot-password [post]
func (auth *AuthHandler) ForgotPassword(c *gin.Context) {
	var req dto.ValidateUserReq
	fmt.Print(req)
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ForgotPassword route"})
}

//		@Summary	 Reset user password
//	 @Description Resets the user's password based on the provided credentials and reset information.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.ResetPasswordReq	  true	"Body"
//		@Success	 200	{object}	response.Response	"Password reset successfully"
//		@Failure	 400	{object}	response.Response	"Bad Request - Invalid input or parameters"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or expired credentials"
//		@Failure	 404	{object}	response.Response	"Not Found - User or resource not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/reset-password [post]
func (auth *AuthHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	err := auth.service.ResetPassword(c, userId, &req)
	if err != nil {
		switch err.Error() {
		case messages.WrongPassword:
			response.Error(c, http.StatusConflict, err, messages.WrongPassword)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	res := dto.ResetPasswordRes{
		Message: "Reset Password Successfully",
	}
	response.JSON(c, http.StatusOK, res)
}
