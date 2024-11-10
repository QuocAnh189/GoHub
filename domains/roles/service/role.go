package service

import (
	"context"
	"gohub/domains/roles/repository"

	"github.com/QuocAnh189/GoBin/validation"
)
type IRoleService interface {
	AddFunction(ctx context.Context)
	RemoveFunction(ctx context.Context)
}

type RoleService struct {
	validations validation.Validation
	repo        repository.IRoleRepository
}

func NewRolService(validations validation.Validation, repo repository.IRoleRepository) *RoleService {
	return &RoleService{
		validations: validations,
		repo:        repo,
	}
}

func (r *RoleService) AddFunction(ctx context.Context) {
	panic("unimplemented")
}

func (r *RoleService) RemoveFunction(ctx context.Context) {
	panic("unimplemented")
}