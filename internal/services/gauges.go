package services

import "github.com/alevnyacow/metrics/internal/domain"

// CountersService provides logic of working with
// gauge metrics.
type GaugesService struct {
	repository GaugesRepository
}

func (service *GaugesService) GetByKey(key domain.GaugeName) (dto domain.Metric, exists bool) {
	if !service.repository.Exists(key) {
		exists = false
		return
	}
	exists = true
	gaugeDTO := service.repository.Get(key)
	dto = gaugeDTO.ToMetricModel()
	return
}

func (service *GaugesService) Set(key domain.GaugeName, value domain.GaugeValue) {
	service.repository.Set(key, value)
}

func (service *GaugesService) GetAll() (metricDTOs []domain.Metric) {
	metricDTOs = make([]domain.Metric, 0)
	gauges := service.repository.GetAll()
	for _, gaugeDTO := range gauges {
		metricDTOs = append(metricDTOs, gaugeDTO.ToMetricModel())
	}
	return
}
