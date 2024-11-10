package service

import (
	"context"
	"gohub/domains/functions/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IFunctionService interface {
	CreateFunction(ctx context.Context)
	GetFunctions(ctx context.Context)
	GetFunction(ctx context.Context)
	UpdateFunction(ctx context.Context)
	DeleteFunction(ctx context.Context)
	EnableCommand(ctx context.Context)
	DisableCommand(ctx context.Context)
}

type FunctionService struct {
	validator 	validation.Validation
	repo 		repository.IFunctionRepository
}

func NewFunctionService(validator validation.Validation, repo repository.IFunctionRepository) *FunctionService {
	return &FunctionService{
		validator: validator,
		repo: repo,
	}
}

func (f *FunctionService) CreateFunction(ctx context.Context) {
    panic("unimplemented")
}

func (f *FunctionService) GetFunctions(ctx context.Context) {
	panic("unimplemented")
}

func (f *FunctionService) GetFunction(ctx context.Context) {
	panic("unimplemented")
}

func (f *FunctionService) UpdateFunction(ctx context.Context) {
	panic("unimplemented")
}

func (f *FunctionService) DeleteFunction(ctx context.Context) {
    panic("unimplemented")
}

func (f *FunctionService) EnableCommand(ctx context.Context) {
	panic("unimplemented")
}

func (f *FunctionService) DisableCommand(ctx context.Context) {
    panic("unimplemented")
}