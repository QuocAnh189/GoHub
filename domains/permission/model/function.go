package model

import (
	"time"
)

type Function struct {
	ID        		string     `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      		string     `json:"name"`
	Url       		string     `json:"url"`
	SortOrder 		string     `json:"sortOrder"`
	ParentId  		string     `json:"parentId"`
	IsDeleted     	bool       `json:"isDeleted" gorm:"default:0"`
	DeletedAt     	*time.Time `json:"deletedAt" gorm:"index"`
	CreatedAt     	time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     	time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
}
