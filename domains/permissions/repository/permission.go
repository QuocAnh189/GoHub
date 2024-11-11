package repository

import "gohub/database"

type IPermissionRepository interface{}

type PermissionRepository struct {
	db database.IDatabase
}

func NewPermissionRepository(db database.IDatabase) *PermissionRepository {
    return &PermissionRepository{db: db}
}