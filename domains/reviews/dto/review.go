package dto

import (
	"gohub/pkg/paging"
)

type Review struct {
	ID         string  `json:"id"`
	User       User    `json:"user"`
	EventId    string  `json:"eventId"`
	Content    string  `json:"content"`
	Rate       float32 `json:"rate"`
	IsPositive bool    `json:"isPositive"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type ReviewByUser struct {
	ID         string  `json:"id"`
	EventId    string  `json:"eventId"`
	Content    string  `json:"content"`
	Rate       float32 `json:"rate"`
	IsPositive bool    `json:"isPositive"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type ReviewByEvent struct {
	ID         string  `json:"id"`
	User       User    `json:"user"`
	EventId    string  `json:"eventId"`
	Content    string  `json:"content"`
	Rate       float32 `json:"rate"`
	IsPositive bool    `json:"isPositive"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type ReviewByCreatedEvent struct {
	ID         string  `json:"id"`
	User       User    `json:"user"`
	Event      Event   `json:"event"`
	Content    string  `json:"content"`
	Rate       float32 `json:"rate"`
	IsPositive bool    `json:"isPositive"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type Event struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	FullName  string `json:"fullName"`
}

type ListReviewReq struct {
	Search    string `json:"content,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"pageSize"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListReviewRes struct {
	Reviews    []*Review          `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type ListReviewByUserRes struct {
	Reviews    []*ReviewByUser    `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type ListReviewByEventRes struct {
	Reviews    []*ReviewByEvent   `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type ListReviewByCreatedEventsRes struct {
	Reviews    []*ReviewByCreatedEvent     `json:"items"`
	Pagination *paging.Pagination          `json:"metadata"`
	Statistic  StatisticReviewCreatedEvent `json:"statistic"`
}

type StatisticReviewCreatedEvent struct {
	AverageRate        float64     `json:"averageRate"`
	TotalPositive      float64     `json:"totalPositive"`
	TotalNegative      float64     `json:"totalNegative"`
	TotalPerNumberRate []RateCount `json:"totalPerNumberRate"`
}

type RateCount struct {
	Rate  int `json:"rate"`
	Total int `json:"value"`
}

type CreateReviewReq struct {
	UserId     string  `form:"userId" validate:"required"`
	EventId    string  `form:"eventId" validate:"required"`
	Content    string  `form:"content" validate:"required"`
	Rate       float32 `form:"rate" validate:"required"`
	IsPositive bool    `form:"isPositive"`
}

type CreateReviewRes struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	Rate      float32 `json:"rate"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type UpdateReviewReq struct {
	ID      string  `form:"id" validate:"required"`
	UserId  string  `form:"userId" validate:"required"`
	EventId string  `form:"eventId" validate:"required"`
	Content string  `form:"content" validate:"required"`
	Rate    float32 `form:"rate" validate:"required"`
}

type UpdateReviewRes struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	Rate      float32 `json:"rate"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
