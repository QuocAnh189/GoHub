package dto

import (
	"gohub/pkg/paging"
)

type Event struct {
	ID                 string      `json:"id"`
	User               *User       `json:"user"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	CoverImageUrl      string      `json:"coverImageUrl"`
	CoverImageFileName string      `json:"coverImageFileName"`
	StartTime          string      `json:"startTime"`
	EndTime            string      `json:"endTime"`
	Location           string      `json:"location"`
	PathLocation       string      `json:"pathLocation"`
	EventCycleType     string      `json:"eventCycleType"`
	EventPaymentType   string      `json:"eventPaymentType"`
	IsPrivate          bool        `json:"isPrivate"`
	SubImage           []*SubImage `json:"subImages"`
	Categories         []*Category `json:"categories"`
	Reasons            []*Reason   `json:"reasons"`
	Coupons            []*Coupon   `json:"coupons"`
	AvgRate            float32     `json:"avgRate"`
}

type Events struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	CoverImageUrl      string  `json:"coverImageUrl"`
	CoverImageFileName string  `json:"coverImageFileName"`
	StartTime          string  `json:"startTime"`
	EndTime            string  `json:"endTime"`
	Location           string  `json:"location"`
	PathLocation       string  `json:"pathLocation"`
	EventCycleType     string  `json:"eventCycleType"`
	EventPaymentType   string  `json:"eventPaymentType"`
	IsPrivate          bool    `json:"isPrivate"`
	AvgRate            float32 `json:"avgRate"`
}

type CreateEventReq struct {
	UserId             string   `form:"userId"`
	Name               string   `form:"name"`
	Description        string   `form:"description"`
	CoverImageUrl      string   `form:"coverImageUrl"`
	CoverImageFileName string   `form:"coverImageFileName"`
	StartTime          string   `form:"startTime"`
	EndTime            string   `form:"endTime"`
	Location           string   `form:"location"`
	PathLocation       string   `form:"pathLocation"`
	EventCycleType     string   `form:"eventCycleType"`
	EventPaymentType   string   `form:"eventPaymentType"`
	IsPrivate          bool     `form:"isPrivate"`
	CategoryIds        []string `form:"categoryIds"`
	ReasonItems        []string `form:"reasonItems"`
}

type UpdateEventReq struct {
	ID                 string   `form:"id"`
	UserId             string   `form:"userId"`
	Name               string   `form:"name"`
	Description        string   `form:"description"`
	CoverImageUrl      string   `form:"coverImageUrl"`
	CoverImageFileName string   `form:"coverImageFileName"`
	StartTime          string   `form:"startTime"`
	EndTime            string   `form:"endTime"`
	Location           string   `form:"location"`
	PathLocation       string   `form:"pathLocation"`
	EventCycleType     string   `form:"eventCycleType"`
	EventPaymentType   string   `form:"eventPaymentType"`
	IsPrivate          bool     `form:"isPrivate"`
	CategoryIds        []string `form:"categoryIds"`
	ReasonItems        []string `form:"reasonItems"`
}

type ListEventReq struct {
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
	IsPrivate bool   `json:"-" form:"is_private"`
}

type ListEventRes struct {
	Events     []*Events          `json:"events"`
	Pagination *paging.Pagination `json:"metadata"`
}

type DeleteRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

type RestoreRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

type CreateEventFavouriteReq struct {
	UserID  string `json:"userId" binding:"required"`
	EventId string `json:"eventId" binding:"required"`
}

type UnEventFavouriteReq struct {
	UserID  string `json:"userId" binding:"required"`
	EventId string `json:"eventId" binding:"required"`
}

type MakeEventPublicOrPrivateReq struct {
	UserId string   `json:"userId"`
	Ids    []string `json:"ids" binding:"required"`
}