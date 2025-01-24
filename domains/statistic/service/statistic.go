package service

import (
	"gohub/domains/statistic/repository"
	"gohub/internal/libs/validation"
)

type IStatisticService interface{}

type StatisticService struct {
	validations   validation.Validation
	statisticRepo repository.IStatisticRepository
}

func NewStatisticService(validations validation.Validation, statisticRepo repository.IStatisticRepository) *StatisticService {
	return &StatisticService{
		validations:   validations,
		statisticRepo: statisticRepo,
	}
}
