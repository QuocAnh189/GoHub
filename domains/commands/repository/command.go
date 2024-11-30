package repository

import (
	"context"
	"gohub/database"
	"gohub/domains/commands/model"
	relation "gohub/domains/shares/model"
)

type ICommandRepository interface {
	GetCommandInFunction(ctx context.Context, functionId string) ([]*model.Command, error)
	GetCommandNotInFunction(ctx context.Context, functionId string) ([]*model.Command, error)
}

type CommandRepo struct {
	db database.IDatabase
}

func NewCommandRepository(db database.IDatabase) *CommandRepo {
	return &CommandRepo{db: db}
}

func (c *CommandRepo) GetCommandInFunction(ctx context.Context, functionId string) ([]*model.Command, error) {
	var commands []*model.Command

	var commandIds []string
	err := c.db.GetDB().Model(&relation.CommandInFunction{}).
		Select("command_id").
		Where("function_id = ?", functionId).
		Find(&commandIds).Error
	if err != nil {
		return nil, err
	}

	query := database.NewQuery("id IN ?", commandIds)
	err = c.db.Find(ctx, &commands, database.WithQuery(query))
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func (c *CommandRepo) GetCommandNotInFunction(ctx context.Context, functionId string) ([]*model.Command, error) {
	var commands []*model.Command

	var commandIds []string
	err := c.db.GetDB().Model(&relation.CommandInFunction{}).
		Select("command_id").
		Where("function_id = ?", functionId).
		Find(&commandIds).Error
	if err != nil {
		return nil, err
	}

	query := database.NewQuery("id NOT IN ?", commandIds)
	err = c.db.Find(ctx, &commands, database.WithQuery(query))
	if err != nil {
		return nil, err
	}

	return commands, nil
}
