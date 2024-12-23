package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/gin-gonic/gin"
	"gohub/domains/payments/dto"
	"gohub/domains/payments/service"
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
