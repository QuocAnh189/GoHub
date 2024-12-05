package dto

import (
	"gohub/pkg/paging"
	"gorm.io/gorm"
	"time"
)

type Conversation struct {
	ID          string  `json:"id"`
	User        User    `json:"user"`
	Organizer   User    `json:"organizer"`
	LastMessage Message `json:"lastMessage"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type ConversationByOrganizer struct {
	ID          string  `json:"id"`
	User        User    `json:"user"`
	LastMessage Message `json:"lastMessage"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type ConversationByUser struct {
	ID          string  `json:"id"`
	Organizer   User    `json:"organizer"`
	LastMessage Message `json:"lastMessage"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	AvatarUrl string `json:"avatarUrl"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
}

type Message struct {
	SenderId   string         `json:"senderId"`
	ReceiverId string         `json:"receiverId"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}

type ListConversationReq struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListConversationByOrganizerRes struct {
	Conversation []*ConversationByOrganizer `json:"conversations"`
	Pagination   *paging.Pagination         `json:"metadata"`
}

type ListConversationByUserRes struct {
	Conversation []*ConversationByUser `json:"conversations"`
	Pagination   *paging.Pagination    `json:"metadata"`
}

type ListMessageReq struct {
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListMessageRes struct {
	Messages   []*Message         `json:"messages"`
	Pagination *paging.Pagination `json:"metadata"`
}

type CreateMessageReq struct {
	ConversationId string `form:"conversationId" validate:"required"`
	SenderId       string `form:"senderId" validate:"required"`
	ReceiverId     string `form:"receiverId" validate:"required"`
	Content        string `form:"content" validate:"required"`
}

type UpdateMessageReq struct {
	ID             string `form:"id" validate:"required"`
	ConversationId string `form:"conversationId" validate:"required"`
	SenderId       string `form:"senderId" validate:"required"`
	ReceiverId     string `form:"receiverId" validate:"required"`
	Content        string `form:"content" validate:"required"`
}
