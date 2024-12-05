package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"gohub/domains/reviews/dto"
	"gohub/domains/reviews/service"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service service.IReviewService
}

func NewReviewHandler(service service.IReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

//		@Summary	 Create a new review
//	 @Description Creates a new review based on the provided details.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Review created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req dto.CreateReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	review, err := h.service.CreateReview(c, &req)
	if err != nil {
		logger.Error("Failed to create review ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create review")
		}
		return
	}

	var res dto.CreateReviewRes
	utils.MapStruct(&res, &review)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a list of reviews
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews [get]
func (h *ReviewHandler) GetReviews(c *gin.Context) {
	var req *dto.ListReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	reviews, pagination, err := h.service.GetReviews(c, req)
	if err != nil {
		logger.Error("Failed to get reviews: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListReviewRes
	utils.MapStruct(&res.Reviews, reviews)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a list of reviews by event
//	 @Description Fetches a paginated list of reviews created by the event, based on the provided pagination filter.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/get-by-event/{eventId} [get]
func (h *ReviewHandler) GetReviewsByEvent(c *gin.Context) {
	var req *dto.ListReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	eventId := c.Param("eventId")
	reviews, pagination, err := h.service.GetReviewsByEvent(c, eventId, req)
	if err != nil {
		logger.Error("Failed to get reviews: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListReviewByEventRes
	utils.MapStruct(&res.Reviews, reviews)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a list of reviews by user
//	 @Description Fetches a paginated list of reviews created by the user, based on the provided pagination filter.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/get-by-user/{userId} [get]
func (h *ReviewHandler) GetReviewsByUser(c *gin.Context) {
	var req *dto.ListReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	userId := c.Param("userId")
	reviews, pagination, err := h.service.GetReviewsByUser(c, userId, req)
	if err != nil {
		logger.Error("Failed to get reviews: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	var res dto.ListReviewByUserRes
	utils.MapStruct(&res.Reviews, reviews)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a reviews by its ID
//	 @Description Fetches the details of a specific review based on the provided review ID.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the review"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/{reviewId} [get]
func (h *ReviewHandler) GetReviewById(c *gin.Context) {
	var res dto.Review

	reviewId := c.Param("id")
	review, err := h.service.GetReviewById(c, reviewId)
	if err != nil {
		logger.Error("Failed to get review detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res, review)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing review
//	 @Description Updates the details of an existing review based on the provided review ID and update information.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the review"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/{reviewId} [patch]
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	reviewId := c.Param("id")
	var req dto.UpdateReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	review, err := h.service.UpdateReview(c, reviewId, &req)
	if err != nil {
		logger.Error("Failed to update category ", err.Error())
		switch err.Error() {
		case messages.CategoryNameExists:
			response.Error(c, http.StatusConflict, err, messages.CategoryNameExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update category")
		}
		return
	}

	var res dto.UpdateReviewRes
	utils.MapStruct(&res, &review)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete a review
//	 @Description Deletes the review with the specified ID.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Review deleted successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/{reviewId} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewId := c.Param("id")

	err := h.service.DeleteReview(c, reviewId)

	if err != nil {
		logger.Error("Failed to delete review: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete review successfully")
}
