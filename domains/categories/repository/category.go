package repository

import "gohub/database"

type ICategoryRepository interface {
}

type CategoryRepository struct {
	db database.IDatabase
}

func NewCategoryRepository(db database.IDatabase) *CategoryRepository {
    return &CategoryRepository{db: db}
}