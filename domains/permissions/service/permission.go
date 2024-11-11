package service

import (
	"context"
	"gohub/domains/permissions/repository"

	"github.com/QuocAnh189/GoBin/validation"
)

type IPermissionService interface {
	GetPermissions(ctx context.Context)
	GetPermissionsByRoles(ctx context.Context, id string)
	GetPermissionsByUser(ctx context.Context, id string)
	
}

type PermissionService struct {
	validator validation.Validation
	repo      repository.IPermissionRepository
}

func NewPermissionService(validator validation.Validation, repo repository.IPermissionRepository) *PermissionService {
	return &PermissionService{
		validator: validator,
		repo:      repo,
	}
}

func (p *PermissionService) GetPermissions(ctx context.Context) {
	panic("unimplemented")
}

func (p *PermissionService) GetPermissionsByRoles(ctx context.Context, id string) {
    panic("unimplemented")
}

func (p *PermissionService) GetPermissionsByUser(ctx context.Context, id string) {
	panic("unimplemented")
}

