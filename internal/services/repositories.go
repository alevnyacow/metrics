package services

import "github.com/alevnyacow/metrics/internal/domain"

type GaugesRepository interface {
	Set(key domain.GaugeName, value domain.GaugeValue)
	Get(key domain.GaugeName) domain.Gauge
	Exists(key domain.GaugeName) bool
	GetAll() []domain.Gauge
}

type CountersRepository interface {
	Set(key domain.CounterName, value domain.CounterValue)
	Get(key domain.CounterName) domain.Counter
	GetValue(key domain.CounterName) domain.CounterValue
	Exists(key domain.CounterName) bool
	GetAll() []domain.Counter
}
