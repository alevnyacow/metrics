package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

type CommonMetricsService interface {
	UpdateMetrics(ctx context.Context, data []domain.Metric) error
}

type InMemoryCommonMetricsService struct {
	countersService *CountersService
	gaugesService   *GaugesService
}

type DBCommonMetricsService struct {
	db                 *sql.DB
	countersRepository CountersRepository
}

func (service *InMemoryCommonMetricsService) UpdateMetrics(ctx context.Context, data []domain.Metric) (err error) {
	for _, metric := range data {
		if metric.IsCounter() {
			value, success := domain.CounterRawValue(metric.Value).ToValue()
			if !success {
				err = errors.New("could not parse counter value")
				log.Err(err).Msg("error on updating counter")
			}
			service.countersService.Update(ctx, domain.CounterName(metric.Name), value)
		}
		if metric.IsGauge() {
			value, success := domain.GaugeRawValue(metric.Value).ToValue()
			if !success {
				err = errors.New("could not parse gauge value")
				log.Err(err).Msg("error on updating gauge")
			}
			service.gaugesService.Set(ctx, domain.GaugeName(metric.Name), value)
		}
	}
	return
}

func (service *DBCommonMetricsService) UpdateMetrics(ctx context.Context, data []domain.Metric) (err error) {
	transaction, transactionError := service.db.BeginTx(ctx, nil)
	defer transaction.Commit()

	if transactionError != nil {
		err = transactionError
		log.Err(err).Msg("Error on creating transaction")
		return
	}
	for _, data := range data {
		if data.IsGauge() {
			value, success := domain.GaugeRawValue(data.Value).ToValue()
			if !success {
				err = errors.New("could not parse gauge value")
				return
			}
			_, updateError := transaction.ExecContext(ctx, "INSERT INTO gauges (name, value) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET value = $2",
				data.Name,
				value,
			)
			if updateError != nil {
				err = updateError
				log.Err(err).Msg("Error on updating gauge")
				return
			}
		}
		if data.IsCounter() {
			value, success := domain.CounterRawValue(data.Value).ToValue()
			if !success {
				err = errors.New("could not parse counter value")
				return
			}
			_, updateError := transaction.ExecContext(ctx, "INSERT INTO counters (name, value) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET value = counters.value + $2",
				data.Name,
				value,
			)
			if updateError != nil {
				err = updateError
				log.Err(err).Msg("Error on creating counter")
				return
			}
		}
	}
	return
}
