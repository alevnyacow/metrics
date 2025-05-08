// Package services contains metrics applications logic grouped by services.
//
// It includes counters service, gauges service and metrics collection service.
package services

import (
	"database/sql"

	"github.com/alevnyacow/metrics/internal/domain"
)

func NewCountersService(repository CountersRepository, afterUpdate func()) *CountersService {
	return &CountersService{
		repository:  repository,
		afterUpdate: afterUpdate,
	}
}

func NewGaugesService(repository GaugesRepository, afterUpdate func()) *GaugesService {
	return &GaugesService{
		repository:  repository,
		afterUpdate: afterUpdate,
	}
}

func NewMetricsCollectionService() *MetricsCollectionService {
	return &MetricsCollectionService{
		gauges:   make([]domain.Gauge, 0),
		counters: make([]domain.Counter, 0),
	}
}

func NewHealtheckService(db *sql.DB) *HealthcheckService {
	return &HealthcheckService{
		db: db,
	}
}
