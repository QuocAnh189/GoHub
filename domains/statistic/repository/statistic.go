package repository

import "gohub/database"

type IStatisticRepository interface {
}

type StatisticRepo struct {
	db database.IDatabase
}

func NewRoleRepository(db database.IDatabase) *StatisticRepo {
	return &StatisticRepo{db: db}
}
