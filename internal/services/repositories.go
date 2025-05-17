package services

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
)

type GaugesRepository interface {
	Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue) error
	Get(ctx context.Context, key domain.GaugeName) (domain.Gauge, error)
	Exists(ctx context.Context, key domain.GaugeName) bool
	GetAll(ctx context.Context) ([]domain.Gauge, error)
}

type CountersRepository interface {
	Set(ctx context.Context, key domain.CounterName, value domain.CounterValue) error
	Get(ctx context.Context, key domain.CounterName) (domain.Counter, error)
	GetValue(ctx context.Context, key domain.CounterName) (domain.CounterValue, error)
	Exists(ctx context.Context, key domain.CounterName) bool
	GetAll(ctx context.Context) ([]domain.Counter, error)
}
