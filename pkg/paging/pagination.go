package paging

import "math"

const (
	DefaultPageSize int64 = 10
)

type Pagination struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Skip        int64 `json:"skip"`
	TotalCount  int64 `json:"totalCount"`
	TotalPages  int64 `json:"totalPages"`
}

func NewPagination(page int64, pageSize int64, total int64) *Pagination {
	var pageInfo Pagination
	limit := DefaultPageSize
	if pageSize > 0 && pageSize <= limit {
		pageInfo.PageSize = pageSize
	} else {
		pageInfo.PageSize = limit
	}

	totalPage := int64(math.Ceil(float64(total) / float64(pageInfo.PageSize)))
	pageInfo.TotalCount = total
	pageInfo.TotalPages = totalPage
	if page < 1 || totalPage == 0 {
		page = 1
	}

	pageInfo.CurrentPage = page
	pageInfo.Skip = (page - 1) * pageInfo.PageSize

	return &pageInfo
}
