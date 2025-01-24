package dto

import (
	"gohub/pkg/paging"
	"mime/multipart"
)

type Event struct {
	ID               string        `json:"id"`
	User             *User         `json:"creator"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	CoverImageUrl    string        `json:"coverImageUrl"`
	StartTime        string        `json:"startTime"`
	EndTime          string        `json:"endTime"`
	Location         string        `json:"location"`
	PathLocation     string        `json:"pathLocation"`
	EventCycleType   string        `json:"eventCycleType"`
	EventPaymentType string        `json:"eventPaymentType"`
	IsPrivate        bool          `json:"isPrivate"`
	SubImage         []*SubImage   `json:"subImages"`
	Categories       []*Category   `json:"categories"`
	Reasons          []*Reason     `json:"reasons"`
	TicketTypes      []*TicketType `json:"ticketTypes"`
	Coupons          []*Coupon     `json:"coupons"`
	AverageRate      float32       `json:"averageRate"`
}

type Events struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	CoverImageUrl      string        `json:"coverImageUrl"`
	CoverImageFileName string        `json:"coverImageFileName"`
	StartTime          string        `json:"startTime"`
	EndTime            string        `json:"endTime"`
	Location           string        `json:"location"`
	PathLocation       string        `json:"pathLocation"`
	EventCycleType     string        `json:"eventCycleType"`
	EventPaymentType   string        `json:"eventPaymentType"`
	IsPrivate          bool          `json:"isPrivate"`
	AverageRate        float32       `json:"averageRate"`
	Categories         []*Category   `json:"categories"`
	TicketTypes        []*TicketType `json:"ticketTypes"`
}

type MyEvent struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	CoverImageUrl string    `json:"coverImageUrl"`
	StartTime     string    `json:"startTime"`
	Location      string    `json:"location"`
	IsPrivate     bool      `json:"isPrivate"`
	DeletedAt     string    `json:"deletedAt"`
	Coupons       []*Coupon `json:"coupons"`
}

type Expense struct {
	ID    string  `json:"id"`
	Total float64 `json:"total"`
}

type MyEventAnalysis struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	CoverImageUrl  string     `json:"coverImageUrl"`
	AverageRate    float32    `json:"averageRate"`
	TotalSale      int        `json:"totalSale"`
	TotalFavourite int        `json:"totalFavourite"`
	Expenses       []*Expense `json:"expenses"`
}

type EventFavourite struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	CoverImageUrl string      `json:"coverImageUrl"`
	StartTime     string      `json:"startTime"`
	AverageRate   float32     `json:"averageRate"`
	Categories    []*Category `json:"categories"`
}

type CreateEventReq struct {
	UserId           string                  `form:"userId"`
	Name             string                  `form:"name"`
	Description      string                  `form:"description"`
	CoverImage       *multipart.FileHeader   `form:"coverImage"`
	SubImageItems    []*multipart.FileHeader `form:"subImageItems"`
	StartTime        string                  `form:"startTime"`
	EndTime          string                  `form:"endTime"`
	Location         string                  `form:"location"`
	EventCycleType   string                  `form:"eventCycleType"`
	EventPaymentType string                  `form:"eventPaymentType"`
	IsPrivate        bool                    `form:"isPrivate"`
	CategoryIds      []string                `form:"categoryIds"`
	TicketTypeItems  []*CreateTicketType     `form:"ticketTypeItems"`
	ReasonItems      []string                `form:"reasonItems"`
}

type UpdateEventReq struct {
	ID                 string                `form:"id"`
	UserId             string                `form:"userId"`
	Name               string                `form:"name"`
	Description        string                `form:"description"`
	CoverImageUrl      string                `form:"coverImageUrl"`
	CoverImageFileName string                `form:"coverImageFileName"`
	CoverImage         *multipart.FileHeader `form:"coverImage"`
	SubImageItems      []interface{}         `form:"subImageItems"`
	StartTime          string                `form:"startTime"`
	EndTime            string                `form:"endTime"`
	Location           string                `form:"location"`
	PathLocation       string                `form:"pathLocation"`
	EventCycleType     string                `form:"eventCycleType"`
	EventPaymentType   string                `form:"eventPaymentType"`
	IsPrivate          bool                  `form:"isPrivate"`
	CategoryIds        []string              `form:"categoryIds"`
	TicketTypeItems    []*CreateTicketType   `form:"ticketTypeItems"`
	ReasonItems        []string              `form:"reasonItems"`
}

type ListEventReq struct {
	Search         string   `json:"name,omitempty" form:"search"`
	CategoryIds    []string `json:"category,omitempty" form:"categoryIds"`
	Status         string   `json:"status,omitempty" form:"status"`
	MinRate        int      `json:"minRate" form:"minRate"`
	Page           int64    `json:"-" form:"page"`
	Limit          int64    `json:"-" form:"pageSize"`
	OrderBy        string   `json:"-" form:"order_by"`
	OrderDesc      bool     `json:"-" form:"orderDesc"`
	TakeAll        bool     `json:"-" form:"take_all"`
	IsPrivate      bool     `json:"-" form:"is_private"`
	Visibility     string   `json:"-" form:"visibility"`
	PaymentType    string   `json:"-" form:"paymentType"`
	StartTimeRange string   `json:"-" form:"startTimeRange"`
	EndTimeRange   string   `json:"-" form:"endTimeRange"`
}

type ListEventRes struct {
	Events     []*Events          `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type ListMyEventRes struct {
	Events     []*MyEvent         `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
	Statistic  StatisticMyEvent   `json:"statistic"`
}

type StatisticMyEvent struct {
	TotalAll     int64 `json:"totalAll"`
	TotalPublic  int64 `json:"totalPublic"`
	TotalPrivate int64 `json:"totalPrivate"`
}

type ListMyEventAnalysisRes struct {
	Events     []*MyEventAnalysis `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type ListEventFavouriteRes struct {
	Events     []*EventFavourite  `json:"items"`
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

type ApplyCouponReq struct {
	Ids []string `json:"ids" binding:"required"`
}

type RemoveCouponReq struct {
	Ids []string `json:"ids" binding:"required"`
}

type UserFavouriteEvent struct {
	UserId  string `json:"userId"`
	EventId string `json:"eventId"`
}
