package repository

import "gohub/database"

type IRoleRepository interface{}

type RoleRepo struct {
	db database.IDatabase
}

func NewRoleRepository(db database.IDatabase) *RoleRepo {
    return &RoleRepo{db: db}
}