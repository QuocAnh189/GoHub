package repository

import (
	"context"
	"gohub/database"
	"gohub/domains/roles/model"
)

type IRoleRepository interface {
	GetRoleByName(ctx context.Context, name string) (*model.Role, error)
}

type RoleRepo struct {
	db database.IDatabase
}

func NewRoleRepository(db database.IDatabase) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) GetRoleByName(ctx context.Context, name string) (*model.Role, error) {
	query := database.NewQuery("name = ?", name)
	var role model.Role

	err := r.db.FindOne(ctx, &role, database.WithQuery(query))
	if err != nil {
		return nil, err
	}

	return &role, nil
}
