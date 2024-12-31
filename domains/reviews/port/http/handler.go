package http

import (
	"context"
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/gin-gonic/gin"
	"gohub/domains/reviews/dto"
	"gohub/domains/reviews/service"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"gohub/proto/gen/pb_reviews"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

type ReviewHandler struct {
	service service.IReviewService
}

func NewReviewHandler(service service.IReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

//		@Summary	 Create a new reviews
//	 @Description Creates a new reviews based on the provided details.
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

	result := Predict(req.Content)
	if result == "Positive" {
		req.IsPositive = true
	}

	review, err := h.service.CreateReview(c, &req)
	if err != nil {
		logger.Error("Failed to create reviews ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create reviews")
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

//		@Summary	 Retrieve a list of reviews by created events
//	 @Description Fetches a paginated list of reviews created by the user, based on the provided pagination filter.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 202	{object}	response.Response	"Successfully retrieved the list of reviews"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/reviews/get-by-created-event [get]
func (h *ReviewHandler) GetReviewsByCreatedEvents(c *gin.Context) {
	var req *dto.ListReviewReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListReviewByCreatedEventsRes
	userId := c.GetString("userId")
	reviews, pagination, err := h.service.GetReviewByCreatedEvents(c, userId, req, &res.Statistic)
	if err != nil {
		logger.Error("Failed to get reviews: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res.Reviews, reviews)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a review by its ID
//	 @Description Fetches the details of a specific reviews based on the provided review ID.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the reviews"
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
		logger.Error("Failed to get reviews detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res, review)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update an existing reviews
//	 @Description Updates the details of an existing reviews based on the provided pb_reviews ID and update information.
//		@Tags		 Reviews
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the reviews"
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
		logger.Error("Failed to update reviews ", err.Error())
		switch err.Error() {
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update reviews")
		}
		return
	}

	var res dto.UpdateReviewRes
	utils.MapStruct(&res, &review)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete a reviews
//	 @Description Deletes the reviews with the specified ID.
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
		logger.Error("Failed to delete reviews: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete reviews successfully")
}

func Predict(reviewContent string) string {
	conn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not connect to server: %v", err)})
		return ""
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb_reviews.NewReviewClient(conn)

	//reviewContent := c.Param("review")

	// Tạo request gRPC
	req := &pb_reviews.ReviewRequest{Content: reviewContent}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, err := client.SentimentAnalysis(ctxTimeout, req)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error while calling Predict: %v", err)})
		return ""
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"review": reviewContent,
	//	"result": resp.Result,
	//})

	return resp.Result
}
