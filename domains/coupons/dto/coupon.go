package dto

import "gohub/pkg/paging"

type Coupon struct {
	ID              string  `json:"id"`
	CoverImageUrl   string  `json:"coverImageUrl" gorm:"not null"`
	Name            string  `json:"name" gorm:"not null"`
	Description     string  `json:"description" gorm:"not null"`
	MinQuantity     int     `json:"minQuantity"`
	MinValue        float64 `json:"minValue"`
	PercentageValue float64 `json:"percentValue"`
	RealValue       float64 `json:"realValue"`
	ExpireDate      string  `json:"expireDate" gorm:"not null"`
}

type ListCouponReq struct {
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListCouponRes struct {
	Coupon     []*Coupon          `json:"coupons"`
	Pagination *paging.Pagination `json:"metadata"`
}

type CreateCouponReq struct {
	UserId          string  `form:"userId"`
	CoverImageUrl   string  `form:"coverImageUrl" validate:"required"`
	Name            string  `form:"name" validate:"required"`
	Description     string  `form:"description" validate:"required"`
	MinQuantity     int     `form:"minQuantity" validate:"required"`
	MinValue        float64 `form:"minValue" validate:"required"`
	PercentageValue float64 `form:"percentValue" validate:"required"`
	RealValue       float64 `form:"realValue" validate:"required"`
	ExpireDate      string  `form:"expireDate" validate:"required"`
}

type UpdateCouponReq struct {
	ID              string  `form:"id" validate:"required"`
	UserId          string  `form:"userId" validate:"required"`
	CoverImageUrl   string  `form:"coverImageUrl" validate:"required"`
	Name            string  `form:"name" validate:"required"`
	Description     string  `form:"description" validate:"required"`
	MinQuantity     int     `form:"minQuantity" validate:"required"`
	MinValue        float64 `form:"minValue" validate:"required"`
	PercentageValue float64 `form:"percentValue" validate:"required"`
	RealValue       float64 `form:"realValue" validate:"required"`
	ExpireDate      string  `form:"expireDate" validate:"required"`
}
