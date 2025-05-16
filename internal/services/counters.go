package services

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

// CountersService provides logic of working with
// counters metrics.
type CountersService struct {
	repository  CountersRepository
	afterUpdate func()
}

func (service *CountersService) Update(ctx context.Context, key domain.CounterName, value domain.CounterValue) error {
	exists, err := service.repository.Exists(ctx, key)
	if err != nil {
		log.Err(err).Msg("Error on obtaining metric existance")
	}
	if !exists {
		service.repository.Set(ctx, key, value)
		log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Created counter")
		return nil
	}
	counterValue, counterValueObtainingError := service.repository.GetValue(ctx, key)
	if counterValueObtainingError != nil {
		log.Err(counterValueObtainingError).Msg("Error on obtain counter value")
		return counterValueObtainingError
	}
	summedValue := value + counterValue
	service.repository.Set(ctx, key, summedValue)
	service.afterUpdate()
	log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Updated counter")
	return nil
}

func (service *CountersService) GetByKey(ctx context.Context, key domain.CounterName) (dto domain.Metric, exists bool, err error) {
	itemExists, errorOnItemExists := service.repository.Exists(ctx, key)
	if errorOnItemExists != nil {
		err = errorOnItemExists
		return
	}
	if !itemExists {
		exists = false
		return
	}
	exists = true
	counterDTO, errorOnObtainingDTO := service.repository.Get(ctx, key)
	if errorOnObtainingDTO != nil {
		err = errorOnObtainingDTO
		return
	}
	dto = counterDTO.ToMetricModel()
	return
}

func (service *CountersService) GetAll(ctx context.Context) (metricDTOs []domain.Metric, err error) {
	metricDTOs = make([]domain.Metric, 0)
	counters, errorOnGetting := service.repository.GetAll(ctx)
	if errorOnGetting != nil {
		err = errorOnGetting
		return
	}
	for _, counterDTO := range counters {
		metricDTOs = append(metricDTOs, counterDTO.ToMetricModel())
	}
	return
}
