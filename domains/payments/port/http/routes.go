package http

import (
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/payments/repository"
	"gohub/domains/payments/service"
	middleware "gohub/pkg/middleware"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	PaymentRepository := repository.NewPaymentRepository(sqlDB)
	PaymentService := service.NewPaymentService(validator, PaymentRepository)
	PaymentHandler := NewPaymentHandler(PaymentService)

	authMiddleware := middleware.JWTAuth()

	expenseRoute := r.Group("/payments").Use(authMiddleware)
	{
		expenseRoute.GET("/get-transactions", PaymentHandler.GetTransactions)
		expenseRoute.GET("/get-orders", PaymentHandler.GetOrders)
		expenseRoute.POST("/checkout", PaymentHandler.Checkout)
	}
}
