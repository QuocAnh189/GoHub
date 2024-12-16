package dto

import "gohub/pkg/paging"

type Coupon struct {
	ID              string  `json:"id"`
	CoverImageUrl   string  `json:"coverImageUrl" gorm:"not null"`
	Name            string  `json:"name" gorm:"not null"`
	Description     string  `json:"description" gorm:"not null"`
	MinQuantity     int     `json:"minQuantity"`
	MinPrice        float64 `json:"minPrice"`
	PercentageValue float64 `json:"percentageValue"`
	ExpireDate      string  `json:"expireDate" gorm:"not null"`
}

type ListCouponReq struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListCouponRes struct {
	Coupon     []*Coupon          `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type CreateCouponReq struct {
	UserId          string  `form:"userId"`
	CoverImageUrl   string  `form:"coverImageUrl" validate:"required"`
	Name            string  `form:"name" validate:"required"`
	Description     string  `form:"description" validate:"required"`
	MinQuantity     int     `form:"minQuantity" validate:"required"`
	MinPrice        int     `form:"minPrice" validate:"required"`
	PercentageValue float64 `form:"percentageValue" validate:"required"`
	ExpireDate      string  `form:"expireDate" validate:"required"`
}

type UpdateCouponReq struct {
	ID              string  `form:"id" validate:"required"`
	CoverImageUrl   string  `form:"coverImageUrl" validate:"required"`
	Name            string  `form:"name" validate:"required"`
	Description     string  `form:"description" validate:"required"`
	MinQuantity     int     `form:"minQuantity" validate:"required"`
	MinPrice        int     `form:"minPrice" validate:"required"`
	PercentageValue float64 `form:"percentageValue" validate:"required"`
	ExpireDate      string  `form:"expireDate" validate:"required"`
}
