package model

import (
	"gorm.io/gorm"
	"time"
)

type Function struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Name      string         `json:"name"`
	Url       string         `json:"url"`
	SortOrder string         `json:"sortOrder"`
	ParentId  string         `json:"parentId"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (Function) TableName() string {
	return "functions"
}
