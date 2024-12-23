package dto

import "gohub/pkg/paging"

type Ticket struct {
	ID            string     `json:"id"`
	TicketNo      string     `json:"ticketNo"`
	CustomerName  string     `json:"customerName"`
	CustomerPhone string     `json:"customerPhone"`
	CustomerEmail string     `json:"customerEmail"`
	Event         Event      `json:"event"`
	TicketType    TicketType `json:"ticketType"`
}

type Event struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type TicketType struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListTicketReq struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListTicketRes struct {
	Ticket     []*Ticket          `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
