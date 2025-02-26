package http

import (
	"github.com/gin-gonic/gin"
	"gohub/configs"
	"gohub/domains/coupons/dto"
	"gohub/domains/coupons/service"
	"gohub/internal/libs/logger"
	"gohub/pkg/messages"
	"gohub/pkg/response"
	"gohub/pkg/utils"
	"net/http"
)

type CouponHandler struct {
	service service.ICouponService
}

func NewCouponHandler(service service.ICouponService) *CouponHandler {
	return &CouponHandler{
		service: service,
	}
}

//		@Summary	 Retrieve a coupons
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons [get]
func (h *CouponHandler) GetCoupons(c *gin.Context) {
	var req dto.ListCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListCouponRes
	coupons, pagination, err := h.service.GetCoupons(c, &req)
	if err != nil {
		logger.Error("Failed to get list coupons: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	utils.MapStruct(&res.Coupon, &coupons)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a coupons by created
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons/get-created-coupons [get]
func (h *CouponHandler) GetCreatedCoupons(c *gin.Context) {
	var req dto.ListCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to parse request query: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userId := c.GetString("userId")
	coupons, pagination, err := h.service.GetCreatedCoupons(c, userId, &req)
	if err != nil {
		logger.Error("Failed to get list coupons: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	var res dto.ListCouponRes
	utils.MapStruct(&res.Coupon, &coupons)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a coupon by its ID
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons/{couponId} [get]
func (h *CouponHandler) GetCouponById(c *gin.Context) {
	var res dto.Coupon

	couponId := c.Param("id")
	coupon, err := h.service.GetCouponById(c, couponId)
	if err != nil {
		logger.Error("Failed to get coupon detail: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	utils.MapStruct(&res, &coupon)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new coupon
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons [post]
func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var req dto.CreateCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	req.UserId = c.GetString("userId")

	stripeKey := configs.GetConfig().StripeSecretKey

	coupon, err := h.service.CreateCoupon(c, &req, stripeKey)
	if err != nil {
		logger.Error("Failed to create coupon ", err.Error())
		switch err.Error() {
		case messages.CouponNameAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.CouponNameAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to create coupon")
		}
		return
	}

	var res dto.Coupon
	utils.MapStruct(&res, &coupon)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Update coupon
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons/{couponId} [put]
func (h *CouponHandler) UpdateCoupon(c *gin.Context) {
	var req dto.UpdateCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	couponId := c.Param("id")
	coupon, err := h.service.UpdateCoupon(c, couponId, &req)
	if err != nil {
		logger.Error("Failed to update coupon ", err.Error())
		switch err.Error() {
		case messages.CouponNameAlreadyExists:
			response.Error(c, http.StatusConflict, err, messages.CouponNameAlreadyExists)
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to update coupon")
		}
		return
	}

	var res dto.Coupon
	utils.MapStruct(&res, &coupon)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete coupon
//	 @Description Fetches the details of a specific category based on the provided category ID.
//		@Tags		 Coupons
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Category created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Category with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/coupons/{couponId} [delete]
func (h *CouponHandler) DeleteCoupon(c *gin.Context) {
	couponId := c.Param("id")

	err := h.service.DeleteCoupon(c, couponId)

	if err != nil {
		logger.Error("Failed to delete coupon: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete coupon successfully")
}
