package http

import (
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/expense/repository"
	"gohub/domains/expense/service"
	"gohub/internal/libs/validation"
	middleware "gohub/pkg/middleware"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	ExpenseRepository := repository.NewExpenseRepository(sqlDB)
	ExpenseService := service.NewExpenseService(validator, ExpenseRepository)
	ExpenseHandler := NewExpenseHandler(ExpenseService)

	authMiddleware := middleware.JWTAuth()

	expenseRoute := r.Group("/expenses").Use(authMiddleware)
	{
		expenseRoute.GET("/get-by-event/:eventId", ExpenseHandler.GetExpensesByEvent)
		expenseRoute.POST("/", ExpenseHandler.CreateExpense)
		expenseRoute.GET("/:id", ExpenseHandler.GetExpenseById)
		expenseRoute.PUT("/:id", ExpenseHandler.UpdateExpense)
		expenseRoute.DELETE("/:id", ExpenseHandler.DeleteExpense)
		expenseRoute.POST("/:id/sub-expense", ExpenseHandler.CreateSubExpense)
		expenseRoute.PUT("/:id/sub-expense/:subExpenseId", ExpenseHandler.UpdateSubExpense)
		expenseRoute.DELETE("/:id/sub-expense/:subExpenseId", ExpenseHandler.DeleteSubExpense)
	}
}
