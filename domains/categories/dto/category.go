package dto

import "gohub/pkg/paging"

type Category struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	IconImageUrl      string `json:"iconImageUrl"`
	IconImageFileName string `json:"iconImageFileName"`
	Color             string `json:"color" `
}

type ListCategoryReq struct {
	Name      string `json:"name,omitempty" form:"name"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListCategoryRes struct {
	Category   []*Category        `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}

type CreateCategoryReq struct {
	Name              string `form:"name" validate:"required"`
	IconImageUrl      string `form:"iconImageUrl" validate:"required"`
	IconImageFileName string `form:"iconImageFileName" validate:"required"`
	Color             string `form:"color" validate:"required"`
}

type UpdateCategoryReq struct {
	ID                string `form:"id" validate:"required"`
	Name              string `form:"name" validate:"required"`
	IconImageUrl      string `form:"iconImageUrl" validate:"required"`
	IconImageFileName string `form:"iconImageFileName" validate:"required"`
	Color             string `form:"color" validate:"required"`
}

type DeleteRequest struct {
	Ids []string `json:"ids" binding:"required"`
}

type RestoreRequest struct {
	Ids []string `json:"ids" binding:"required"`
}
