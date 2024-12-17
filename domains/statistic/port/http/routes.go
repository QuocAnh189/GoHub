package http

import (
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"
	"gohub/database"
	"gohub/domains/statistic/repository"
	"gohub/domains/statistic/service"
	middleware "gohub/pkg/middlewares"
)

func Routes(r *gin.RouterGroup, sqlDB database.IDatabase, validator validation.Validation) {
	statisticRepository := repository.IStatisticRepository(sqlDB)
	statisticService := service.NewStatisticService(validator, statisticRepository)
	statisticHandler := NewStatisticHandler(statisticService)

	authMiddleware := middleware.JWTAuth()

	statisticRoute := r.Group("/statistic").Use(authMiddleware)
	{
		statisticRoute.GET("/customer-retention-rate", statisticHandler.CustomerRetentionRate)
		statisticRoute.GET("/customer-conversion-rate", statisticHandler.CustomerConversionRate)
		statisticRoute.GET("/customer-segmentation-age", statisticHandler.CustomerSegmentationRate)
		statisticRoute.GET("/event-total-statistic", statisticHandler.EventTotalStatistic)
		statisticRoute.GET("/event-period-revenue", statisticHandler.EventPeriodRevenue)
		statisticRoute.GET("/event-review-by-customer", statisticHandler.EventRevenueByCustomer)
		statisticRoute.GET("/overview-total-statistic", statisticHandler.OverviewTotalStatistic)
		statisticRoute.GET("/overview-sale-statistic", statisticHandler.OverviewSaleStatistic)
	}
}
