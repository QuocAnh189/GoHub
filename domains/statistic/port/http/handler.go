package http

import (
	"github.com/gin-gonic/gin"
	"gohub/domains/statistic/service"
	"gohub/pkg/response"
	"net/http"
)

type StatisticHandler struct {
	service service.IStatisticService
}

func NewStatisticHandler(service service.IStatisticService) *StatisticHandler {
	return &StatisticHandler{service: service}
}

func (h *StatisticHandler) CustomerRetentionRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerRetentionRate route"})
}

func (h *StatisticHandler) CustomerConversionRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerConversionRate route"})
}

func (h *StatisticHandler) CustomerSegmentationRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerSegmentationRate route"})
}

func (h *StatisticHandler) EventTotalStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventTotalStatistic route"})
}

func (h *StatisticHandler) EventPeriodRevenue(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventPeriodRevenue route"})
}

func (h *StatisticHandler) EventRevenueByCustomer(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventRevenueByCustomer route"})
}

func (h *StatisticHandler) OverviewTotalStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test OverviewTotalStatistic route"})
}

func (h *StatisticHandler) OverviewSaleStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test OverviewSaleStatistic route"})
}
