package repository

import "gohub/database"

type ICommandRepository interface{}

type CommandRepo struct {
	db database.IDatabase
}

func NewCommandRepository(db database.IDatabase) *CommandRepo {
    return &CommandRepo{db: db}
}