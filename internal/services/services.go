// Package services contains metrics applications logic grouped by services.
//
// It includes counters service, gauges service and metrics collector service.
package services

import "github.com/alevnyacow/metrics/internal/domain"

func NewCountersService(repository CountersRepository) *CountersService {
	return &CountersService{
		repository: repository,
	}
}

func NewGaugesService(repository GaugesRepository) *GaugesService {
	return &GaugesService{
		repository: repository,
	}
}

func NewMetricsCollectorService() *MetricsCollectorService {
	return &MetricsCollectorService{
		Gauges:   make([]domain.Gauge, 0),
		Counters: make([]domain.Counter, 0),
	}
}
