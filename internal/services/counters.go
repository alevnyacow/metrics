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

func (service *CountersService) Update(ctx context.Context, key domain.CounterName, value domain.CounterValue) {
	if !service.repository.Exists(ctx, key) {
		service.repository.Set(ctx, key, value)
		log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Created counter")
		return
	}
	summedValue := value + service.repository.GetValue(ctx, key)
	service.repository.Set(ctx, key, summedValue)
	service.afterUpdate()
	log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Updated counter")
}

func (service *CountersService) GetByKey(ctx context.Context, key domain.CounterName) (dto domain.Metric, exists bool) {
	if !service.repository.Exists(ctx, key) {
		exists = false
		return
	}
	exists = true
	counterDTO := service.repository.Get(ctx, key)
	dto = counterDTO.ToMetricModel()
	return
}

func (service *CountersService) GetAll(ctx context.Context) (metricDTOs []domain.Metric) {
	metricDTOs = make([]domain.Metric, 0)
	counters := service.repository.GetAll(ctx)
	for _, counterDTO := range counters {
		metricDTOs = append(metricDTOs, counterDTO.ToMetricModel())
	}
	return
}
