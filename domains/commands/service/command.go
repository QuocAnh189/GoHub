package service

import (
	"context"
	"gohub/domains/commands/model"
	"gohub/domains/commands/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type ICommandService interface {
	GetInFunction(ctx context.Context, functionId string) ([]*model.Command, error)
	GetNotInFunction(ctx context.Context, functionId string) ([]*model.Command, error)
}

type CommandService struct {
	validator validation.Validation
	repo      repository.ICommandRepository
}

func NewCommandService(validator validation.Validation, repo repository.ICommandRepository) *CommandService {
	return &CommandService{validator: validator, repo: repo}
}

func (c *CommandService) GetInFunction(ctx context.Context, functionId string) ([]*model.Command, error) {
	commands, err := c.repo.GetCommandInFunction(ctx, functionId)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func (c *CommandService) GetNotInFunction(ctx context.Context, functionId string) ([]*model.Command, error) {
	commands, err := c.repo.GetCommandNotInFunction(ctx, functionId)
	if err != nil {
		return nil, err
	}

	return commands, nil
}
