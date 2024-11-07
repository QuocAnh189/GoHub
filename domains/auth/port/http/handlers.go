package http

import (
	"fmt"
	"gohub/domains/auth/dto"
	"gohub/domains/auth/service"
	"gohub/pkg/response"
	"net/http"

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

//	@Summary	 Validate user credentials
//  @Description Validate user information session one when signup
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.ValidateUserReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"User credentials are valid"
//	@Failure	 400	{object}	response.Response	"Invalid user credentials or request data"
//	@Failure	 404	{object}	response.Response	"Not Found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/validate-user [post]
func (auth *AuthHandler) ValidateUser(c *gin.Context) {
	var req dto.ValidateUserReq
	fmt.Print(req)
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ValidateUser route"})
}

//	@Summary	 Signup a new user
//  @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.SignUpReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"User successfully registered"
//	@Failure	 400	{object}	response.Response	"Invalid user input"
//	@Failure	 404	{object}	response.Response	"Not Found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/signup [post]
func (auth *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpReq
	fmt.Print(req)
	response.JSON(c, http.StatusOK, gin.H{"message": "Test SignUp route"})
}

//	@Summary	 Signin a user
//  @Description Authenticates the user based on the provided credentials and returns a sign-in response if successful.
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.SignInReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"Successfully signed in"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Invalid credentials"
//	@Failure	 404	{object}	response.Response	"Not Found - User not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/signin [post]
func (auth *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInReq
	fmt.Print(req)
	response.JSON(c, http.StatusOK, gin.H{"message": "Test SignIn route"})
}

//	@Summary	 Signout a new user
//  @Description Signs out the current user, invalidating their session or authentication token.
//	@Tags		 Auth
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully signed out"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/signout [post]
func (auth *AuthHandler) SignOut(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test SignIn route"})
}

//	@Summary	 Signin a new user with external credentials
//  @Description Authenticates the user using an external authentication provider (e.g., Google, Facebook) and returns a login response if successful.
//	@Tags		 Auth
//	@Produce	 json
// 	@Param 	     provider query string true "Authentication provider" 
// 	@Param       returnUrl query string true "Redirect URL after authentication" 
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/external-login [post]
func (auth *AuthHandler) ExternalSignIn(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ExternalSignIn route"})
}

//	@Summary	 Callback endpoint for external authentication
//  @Description Handles the callback from an external authentication provider and processes the authentication result.
//	@Tags		 Auth
//	@Produce	 json
// 	@Param       returnUrl query string true "Redirect URL after authentication" 
//	@Failure	 400	{object}	response.Response	"Invalid user credentials or request data"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/external-auth-callback [get]
func (auth *AuthHandler) ExternalCallback(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ExternalCallback route"})
}

//	@Summary	 Refresh user authentication token
//  @Description Refreshes the user's authentication token based on the provided refresh token credentials.
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.RefreshTokenReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"Successfully refreshed the token"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or expired refresh token"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/refresh-token [post]
func (auth *AuthHandler) RefreshToken(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test RefreshToken route"})
}

//	@Summary	 Initiate password recovery
//  @Description Sends a password recovery email to the user based on the provided email address.
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.ForgotPasswordReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"Password recovery email sent successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or missing credentials"
//	@Failure	 404	{object}	response.Response	"Not Found - User with the provided email address not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/forgot-password [post]
func (auth *AuthHandler) ForgotPassword(c *gin.Context) {
	var req dto.ValidateUserReq
	fmt.Print(req)
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ForgotPassword route"})
}

//	@Summary	 Reset user password
//  @Description Resets the user's password based on the provided credentials and reset information.
//	@Tags		 Auth
//	@Produce	 json
//	@Param		 _	body	dto.ResetPasswordReq	  true	"Body"
//	@Success	 200	{object}	response.Response	"Password reset successfully"
//	@Failure	 400	{object}	response.Response	"Bad Request - Invalid input or parameters"
//	@Failure	 401	{object}	response.Response	"Unauthorized - Invalid or expired credentials"
//	@Failure	 404	{object}	response.Response	"Not Found - User or resource not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/reset-password [post]
func (auth *AuthHandler) ResetPassword(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test ResetPassword route"})
}

//	@Summary	 Retrieve user profile
//  @Description Fetches the details of the currently authenticated user.
//	@Tags		 Auth
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"User profile retrieved successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/auth/profile [get]
func (auth *AuthHandler) GetProfile(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test Profile route"})
}
