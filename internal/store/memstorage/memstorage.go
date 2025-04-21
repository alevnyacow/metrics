// Package memstorage contains in-memory repository implementations.
package memstorage

import "github.com/alevnyacow/metrics/internal/domain"

func NewCountersRepository() *CountersRepository {
	return &CountersRepository{
		data: make(map[domain.CounterName]domain.CounterValue),
	}
}

func NewGaugesRepository() *GaugesRepository {
	return &GaugesRepository{
		data: make(map[domain.GaugeName]domain.GaugeValue),
	}
}
