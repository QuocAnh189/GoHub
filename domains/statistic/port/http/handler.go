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

//		@Summary	 Retrieve the customer retention rate
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the customer retention rate"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/customer-retention-rate [get]
func (h *StatisticHandler) CustomerRetentionRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerRetentionRate route"})
}

//		@Summary	 Retrieve the customer conversion rate
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the customer conversion rate"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/customer-conversion-rate [get]
func (h *StatisticHandler) CustomerConversionRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerConversionRate route"})
}

//		@Summary	 Retrieve the customer segmentation rate
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the customer segmentation rate"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/customer-segmentation-rate [get]
func (h *StatisticHandler) CustomerSegmentationRate(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test CustomerSegmentationRate route"})
}

//		@Summary	 Retrieve the event total statistic
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the event total statistic"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/event-total-statistic [get]
func (h *StatisticHandler) EventTotalStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventTotalStatistic route"})
}

//		@Summary	 Retrieve the event period revenue
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the Event Period Revenue"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/event-period-revenue [get]
func (h *StatisticHandler) EventPeriodRevenue(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventPeriodRevenue route"})
}

//		@Summary	 Retrieve the event reviews by customer
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the event reviews by customer"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/event-reviews-by-customer [get]
func (h *StatisticHandler) EventRevenueByCustomer(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test EventRevenueByCustomer route"})
}

//		@Summary	 Retrieve the overview total statistic
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the overview total statistic"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/overview-total-statistic [get]
func (h *StatisticHandler) OverviewTotalStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test OverviewTotalStatistic route"})
}

//		@Summary	 Retrieve the overview sale statistic
//	 @Description Fetches a paginated list of reviews based on the provided filter parameters.
//		@Tags		 Statistics
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the overview sale statistic"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/statistics/overview-sale-statistic [get]
func (h *StatisticHandler) OverviewSaleStatistic(c *gin.Context) {
	response.JSON(c, http.StatusOK, gin.H{"message": "Test OverviewSaleStatistic route"})
}
