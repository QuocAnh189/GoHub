package dto

import "gohub/pkg/paging"

type Transaction struct {
	ID             string  `json:"id"`
	Event          Event   `json:"event"`
	CustomerName   string  `json:"customerName"`
	TicketQuantity int     `json:"ticketQuantity"`
	TotalPrice     float32 `json:"totalPrice"`
	DiscountPrice  float32 `json:"discountPrice"`
	FinalPrice     float32 `json:"finalPrice"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"createdAt"`
}

type Event struct {
	Name          string `json:"name"`
	CoverImageURL string `json:"coverImageUrl"`
}

type ListTransactionReq struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
	StartDate string `json:"-" form:"startDate"`
	EndDate   string `json:"-" form:"endDate"`
}

type ListTransactionRes struct {
	Transaction []*Transaction     `json:"items"`
	Pagination  *paging.Pagination `json:"metadata"`
}

type Order struct {
	ID             string     `json:"id"`
	Event          EventOrder `json:"event"`
	TicketQuantity int        `json:"ticketQuantity"`
	TotalPrice     float32    `json:"totalPrice"`
	DiscountPrice  float32    `json:"discountPrice"`
	FinalPrice     float32    `json:"finalPrice"`
	Status         string     `json:"status"`
	CreatedAt      string     `json:"createdAt"`
}

type EventOrder struct {
	Name          string `json:"name"`
	CoverImageURL string `json:"coverImageUrl"`
	User          User   `json:"creator"`
}

type User struct {
	UserName  string `json:"userName"`
	AvatarUrl string `json:"avatarUrl"`
}

type ListOrderReq struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
	StartDate string `json:"-" form:"startDate"`
	EndDate   string `json:"-" form:"endDate"`
}

type ListOrderRes struct {
	Order      []*Order           `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
