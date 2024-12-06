package http

import (
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/gin-gonic/gin"
	"gohub/domains/coupons/dto"
	"gohub/domains/coupons/service"
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

func (h *CouponHandler) CreateCoupon(c *gin.Context) {
	var req dto.CreateCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	req.UserId = c.GetString("userId")

	coupon, err := h.service.CreateCoupon(c, &req)
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

func (h *CouponHandler) UpdateCoupon(c *gin.Context) {
	couponId := c.Param("id")
	var req dto.UpdateCouponReq
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

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
