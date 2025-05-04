package services

import (
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

// CountersService provides logic of working with
// counters metrics.
type CountersService struct {
	repository  CountersRepository
	afterUpdate func()
}

func (service *CountersService) Update(key domain.CounterName, value domain.CounterValue) {
	if !service.repository.Exists(key) {
		service.repository.Set(key, value)
		log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Created counter")
		return
	}
	summedValue := value + service.repository.GetValue(key)
	service.repository.Set(key, summedValue)
	service.afterUpdate()
	log.Info().Str("Counter name", string(key)).Str("Counter value", value.ToString()).Msg("Updated counter")
}

func (service *CountersService) GetByKey(key domain.CounterName) (dto domain.Metric, exists bool) {
	if !service.repository.Exists(key) {
		exists = false
		return
	}
	exists = true
	counterDTO := service.repository.Get(key)
	dto = counterDTO.ToMetricModel()
	return
}

func (service *CountersService) GetAll() (metricDTOs []domain.Metric) {
	metricDTOs = make([]domain.Metric, 0)
	counters := service.repository.GetAll()
	for _, counterDTO := range counters {
		metricDTOs = append(metricDTOs, counterDTO.ToMetricModel())
	}
	return
}
