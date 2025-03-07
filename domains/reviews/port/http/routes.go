package http

import (
	"gohub/database"
	"gohub/domains/reviews/repository"
	"gohub/domains/reviews/service"
	middleware "gohub/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gohub/internal/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	reviewRepository := repository.NewReviewRepository(sqlDB)
	reviewService := service.NewReviewService(validator, reviewRepository)
	reviewHandler := NewReviewHandler(reviewService)

	authMiddleware := middleware.JWTAuth()
	reviewRoute := r.Group("/reviews").Use(authMiddleware)
	{
		reviewRoute.POST("/", reviewHandler.CreateReview)
		reviewRoute.GET("/", reviewHandler.GetReviews)
		reviewRoute.GET("/:id", reviewHandler.GetReviewById)
		reviewRoute.GET("/get-by-event/:eventId", reviewHandler.GetReviewsByEvent)
		reviewRoute.GET("/get-by-user/:userId", reviewHandler.GetReviewsByUser)
		reviewRoute.GET("/get-by-created-events", reviewHandler.GetReviewsByCreatedEvents)
		reviewRoute.PUT("/:id", reviewHandler.UpdateReview)
		reviewRoute.DELETE("/:id", reviewHandler.DeleteReview)
	}
}
