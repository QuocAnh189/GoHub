package service

import (
	"context"
	"gohub/domains/commands/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type ICommandService interface{
	GetInFunction(ctx context.Context)
	GetNotInFunction(ctx context.Context)
}

type CommandService struct {
	validator validation.Validation
	repo repository.ICommandRepository
}

func NewCommandService(validator validation.Validation, repo repository.ICommandRepository) *CommandService {
    return &CommandService{validator: validator, repo: repo}
}

func (c *CommandService) GetInFunction(ctx context.Context) {
	panic("unimplemented")
}

func (c *CommandService) GetNotInFunction(ctx context.Context) {
	panic("unimplemented")
}

