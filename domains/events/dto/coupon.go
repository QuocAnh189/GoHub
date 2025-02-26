package dto

import "time"

type Coupon struct {
	ID              string    `json:"id"`
	CouponId        string    `json:"couponId"`
	Name            string    `json:"name"`
	CoverImageUrl   string    `json:"coverImageUrl"`
	Description     string    `json:"description"`
	MinQuantity     int       `json:"minQuantity"`
	MinPrice        float64   `json:"minPrice"`
	PercentageValue float64   `json:"percentageValue"`
	ExpireDate      time.Time `json:"expireDate"`
}
