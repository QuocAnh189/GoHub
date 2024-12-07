package dto

import (
	"gohub/pkg/paging"
)

type Expense struct {
	ID          string        `json:"id"`
	EventId     string        `json:"eventId"`
	Title       string        `json:"title"`
	Total       float64       `json:"total"`
	SubExpenses *[]SubExpense `json:"subExpenses"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
}

type ListExpenseReq struct {
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListExpenseRes struct {
	Expense    []*Expense         `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type CreatedExpenseReq struct {
	Title   string  `json:"title"`
	EventId string  `json:"eventId"`
	Total   float64 `json:"total"`
}

type UpdatedExpenseReq struct {
	ID      string  `json:"id"`
	EventId string  `json:"eventId"`
	Title   string  `json:"title"`
	Total   float64 `json:"total"`
}
