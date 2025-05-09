package services

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

// CountersService provides logic of working with
// gauge metrics.
type GaugesService struct {
	repository  GaugesRepository
	afterUpdate func()
}

func (service *GaugesService) GetByKey(ctx context.Context, key domain.GaugeName) (dto domain.Metric, exists bool) {
	if !service.repository.Exists(ctx, key) {
		exists = false
		return
	}
	exists = true
	gaugeDTO := service.repository.Get(ctx, key)
	dto = gaugeDTO.ToMetricModel()
	return
}

func (service *GaugesService) Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue) {
	service.repository.Set(ctx, key, value)
	service.afterUpdate()
	log.Info().Str("Gauge name", string(key)).Str("Gauge value", value.ToString()).Msg("Setted gauge")

}

func (service *GaugesService) GetAll(ctx context.Context) (metricDTOs []domain.Metric) {
	metricDTOs = make([]domain.Metric, 0)
	gauges := service.repository.GetAll(ctx)
	for _, gaugeDTO := range gauges {
		metricDTOs = append(metricDTOs, gaugeDTO.ToMetricModel())
	}
	return
}
