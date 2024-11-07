package http

import (
	"gohub/database"
	"gohub/domains/reviews/repository"
	"gohub/domains/reviews/service"

	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	reviewRepository := repository.NewReviewRepository(sqlDB)
	reviewService := service.NewReviewService(validator, reviewRepository)
	reviewHandler := NewReviewHandler(reviewService)

	reviewRoute := r.Group("/reviews")
	{
		reviewRoute.POST("/", reviewHandler.CreateReview)
		reviewRoute.GET("/", reviewHandler.GetReviews)
		reviewRoute.GET("/:id", reviewHandler.GetReview)
		reviewRoute.GET("/get-by-event/:eventId", reviewHandler.GetReview)
		reviewRoute.GET("/get-by-user/:userId", reviewHandler.GetReview)
		reviewRoute.PUT("/:id", reviewHandler.UpdateReview)
		reviewRoute.DELETE("/:id", reviewHandler.DeleteReview)
	}
}