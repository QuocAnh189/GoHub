package http

import (
	"gohub/domains/reviews/service"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service service.IReviewService
}

func NewReviewHandler(service service.IReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}


//	@Summary	 Create a new review
//  @Description Creates a new review based on the provided details.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 201	{object}	response.Response	"Review created successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CreateReview route"})
}

//	@Summary	 Retrieve a list of reviews
//  @Description Fetches a paginated list of reviews based on the provided filter parameters.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews [get]
func (h *ReviewHandler) GetReviews(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetReviews route"})
}

//	@Summary	 Retrieve a list of reviews by event
//  @Description Fetches a paginated list of reviews created by the event, based on the provided pagination filter.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews/get-by-event/{eventId} [get]
func (h *ReviewHandler) GetReviewsByEvent(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetReviewByEvent route"})
}

//	@Summary	 Retrieve a list of reviews by user
//  @Description Fetches a paginated list of reviews created by the user, based on the provided pagination filter.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews/get-by-user/{userId} [get]
func (h *ReviewHandler) GetReviewsByUser(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetReviewByUser route"})
}

//	@Summary	 Retrieve a reviews by its ID
//  @Description Fetches the details of a specific review based on the provided review ID.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the review"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews/{reviewId} [get]
func (h *ReviewHandler) GetReview(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test GetReview route"})
}

//	@Summary	 Update an existing review
//  @Description Updates the details of an existing review based on the provided review ID and update information.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Successfully retrieved the review"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews/{reviewId} [patch]
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test UpdateReview route"})
}

//	@Summary	 Delete a review
//  @Description Deletes the review with the specified ID.
//	@Tags		 Reviews
//	@Produce	 json
//	@Success	 200	{object}	response.Response	"Review deleted successfully"
//	@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//	@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router		 /api/v1/reviews/{reviewId} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{"message": "Test DeleteReview route"})
}


