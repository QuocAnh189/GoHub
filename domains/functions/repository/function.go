package repository

import "gohub/database"

type IFunctionRepository interface{}

type FunctionRepository struct {
	db database.IDatabase
}

func NewFunctionRepository(db database.IDatabase) *FunctionRepository {
    return &FunctionRepository{db: db}
}