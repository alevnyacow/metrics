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

func (service *GaugesService) GetByKey(ctx context.Context, key domain.GaugeName) (dto domain.Metric, exists bool, err error) {
	metricExists := service.repository.Exists(ctx, key)
	if !metricExists {
		exists = false
		return
	}
	exists = true
	gaugeDTO, errOnGettingDTO := service.repository.Get(ctx, key)
	if errOnGettingDTO != nil {
		err = errOnGettingDTO
		return
	}
	dto = gaugeDTO.ToMetricModel()
	return
}

func (service *GaugesService) Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue) error {
	errOnSet := service.repository.Set(ctx, key, value)
	if errOnSet != nil {
		return errOnSet
	}
	service.afterUpdate()
	log.Info().Str("Gauge name", string(key)).Str("Gauge value", value.ToString()).Msg("Setted gauge")
	return nil
}

func (service *GaugesService) GetAll(ctx context.Context) (metricDTOs []domain.Metric, err error) {
	metricDTOs = make([]domain.Metric, 0)
	gauges, errOnGettingAll := service.repository.GetAll(ctx)
	if errOnGettingAll != nil {
		err = errOnGettingAll
		return
	}
	for _, gaugeDTO := range gauges {
		metricDTOs = append(metricDTOs, gaugeDTO.ToMetricModel())
	}
	return
}
