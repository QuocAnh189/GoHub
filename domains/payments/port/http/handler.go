package http

import (
	"github.com/gin-gonic/gin"
	"gohub/configs"
	"gohub/domains/payments/dto"
	"gohub/domains/payments/service"
	"gohub/internal/libs/logger"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"
)

type PaymentHandler struct {
	service service.IPaymentService
}

func NewPaymentHandler(service service.IPaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

//		@Summary	 Retrieve a transactions
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Payments
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/payments/get-transactions [get]
func (h *PaymentHandler) GetTransactions(c *gin.Context) {
	var req dto.ListTransactionReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	transactions, pagination, err := h.service.GetTransactions(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get list transactions: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListTransactionRes
	utils.MapStruct(&res.Transaction, &transactions)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a orders
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Payments
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/payments/get-orders [get]
func (h *PaymentHandler) GetOrders(c *gin.Context) {
	var req dto.ListOrderReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	orders, pagination, err := h.service.GetOrders(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get list transactions: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListOrderRes
	utils.MapStruct(&res.Order, &orders)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Checkout
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Payments
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/payments/create-session [post]
func (h *PaymentHandler) CreateSession(c *gin.Context) {
	cfg := configs.GetConfig()

	var req dto.TicketCheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	sessionId, sessionUrl, paymentId, err := h.service.CreateSession(c, &req, cfg.StripeSecretKey)

	if err != nil {
		logger.Error("Failed to checkout: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	result := dto.TicketCheckoutResponse{
		SessionID:  sessionId,
		SessionUrl: sessionUrl,
		PaymentId:  paymentId,
		Data:       req,
	}
	response.JSON(c, http.StatusOK, result)
}

//		@Summary	 Checkout
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Payments
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/payments/checkout [post]
func (h *PaymentHandler) Checkout(c *gin.Context) {
	var req dto.TicketCheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.service.Checkout(c, &req); err != nil {
		logger.Error("Failed to checkout: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Some thing went wrong")
	}

	response.JSON(c, http.StatusOK, true)
}
