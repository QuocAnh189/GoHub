package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"gohub/configs"
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
//		@Router		 /api/v1/payments/checkout [post]
func (h *PaymentHandler) Checkout(c *gin.Context) {
	cfg := configs.GetConfig()
	stripe.Key = cfg.StripeSecretKey

	var req dto.TicketPurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams
	for _, item := range req.TicketItems {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"), // Thay đổi thành đơn vị tiền tệ của bạn
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: item.Type,
				},
				UnitAmount: stripe.Int64(int64(item.Price)), // Giá vé (đơn vị: cents)
			},
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	// Tạo một Checkout Session trong Stripe
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String("http://localhost:3000"), // Thay bằng URL thành công của bạn
		CancelURL:          stripe.String("http://localhost:3000"), // Thay bằng URL hủy của bạn
		CustomerEmail:      stripe.String(req.Email),
	}

	s, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//result := dto.TicketPurchaseResponse{
	//	SessionID:   s.ID,
	//	CheckoutURL: s.URL,
	//}
	//response.JSON(c, http.StatusOK, result)

	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Redirect(http.StatusFound, s.URL)

	//response.JSON(c, http.StatusFound, result.CheckoutURL)
}
