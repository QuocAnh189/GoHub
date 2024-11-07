package repository

import "gohub/database"

type IConversationRepository interface {
}

type ConversationRepo struct {
	db database.IDatabase
}

func NewConversationRepository(db database.IDatabase) *ConversationRepo {
	return &ConversationRepo{db: db}
}