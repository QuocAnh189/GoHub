package service

import (
	"github.com/QuocAnh189/GoBin/validation"
	"gohub/domains/statistic/repository"
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
