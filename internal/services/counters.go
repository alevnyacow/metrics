package services

import "github.com/alevnyacow/metrics/internal/domain"

// CountersService provides logic of working with
// counters metrics.
type CountersService struct {
	repository CountersRepository
}

func (service *CountersService) SetWithRawValue(key domain.CounterName, rawValue domain.CounterRawValue) (success bool) {
	value, parsed := rawValue.ToValue()
	if !parsed {
		success = false
		return
	}
	success = true
	if !service.repository.Exists(key) {
		service.repository.Set(key, value)
		return
	}
	summedValue := value + service.repository.GetValue(key)
	service.repository.Set(key, summedValue)
	return
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
