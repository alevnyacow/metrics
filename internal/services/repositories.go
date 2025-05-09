package services

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
)

type GaugesRepository interface {
	Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue)
	Get(ctx context.Context, key domain.GaugeName) domain.Gauge
	Exists(ctx context.Context, key domain.GaugeName) bool
	GetAll(ctx context.Context) []domain.Gauge
}

type CountersRepository interface {
	Set(ctx context.Context, key domain.CounterName, value domain.CounterValue)
	Get(ctx context.Context, key domain.CounterName) domain.Counter
	GetValue(ctx context.Context, key domain.CounterName) domain.CounterValue
	Exists(ctx context.Context, key domain.CounterName) bool
	GetAll(ctx context.Context) []domain.Counter
}
