package dto

import "time"

type Coupon struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	MinQuantity     int       `json:"minQuantity"`
	MinValue        float64   `json:"minValue"`
	PercentageValue float64   `json:"percentValue"`
	RealValue       float64   `json:"realValue"`
	ExpireDate      time.Time `json:"expireDate"`
}
