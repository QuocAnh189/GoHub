package http

import (
	"gohub/database"
	"gohub/domains/coupons/repository"
	"gohub/domains/coupons/service"
	middleware "gohub/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gohub/internal/libs/validation"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	CouponRepository := repository.NewCouponRepository(sqlDB)
	CouponService := service.NewCouponService(validator, CouponRepository)
	CouponHandler := NewCouponHandler(CouponService)

	authMiddleware := middleware.JWTAuth()

	categoryRoute := r.Group("/coupons").Use(authMiddleware)
	{
		categoryRoute.GET("/", CouponHandler.GetCoupons)
		categoryRoute.GET("/get-created-coupons", CouponHandler.GetCreatedCoupons)
		categoryRoute.GET("/:id", CouponHandler.GetCouponById)
		categoryRoute.POST("/", CouponHandler.CreateCoupon)
		categoryRoute.PUT("/:id", CouponHandler.UpdateCoupon)
		categoryRoute.DELETE("/:id", CouponHandler.DeleteCoupon)
	}
}
