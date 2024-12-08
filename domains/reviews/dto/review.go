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
	Content    string  `json:"content"`
	Rate       float32 `json:"rate"`
	IsPositive bool    `json:"isPositive"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	FullName  string `json:"fullName"`
}

type ListReviewReq struct {
	Content   string `json:"content,omitempty" form:"content"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
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

type CreateReviewReq struct {
	UserId  string  `form:"userId" validate:"required"`
	EventId string  `form:"eventId" validate:"required"`
	Content string  `form:"content" validate:"required"`
	Rate    float32 `form:"rate" validate:"required"`
}

type CreateReviewRes struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	Rate      float32 `json:"rate"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type UpdateReviewReq struct {
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
