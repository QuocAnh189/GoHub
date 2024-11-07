package repository

import "gohub/database"

type IEventRepository interface {
}

type EventRepo struct {
	db database.IDatabase
}

func NewEventRepository(db database.IDatabase) *EventRepo {
	return &EventRepo{db: db}
}