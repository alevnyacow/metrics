package memstorage

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
)

// GaugesRepository is in-memory implementation
// of GaugesRepository interface.
type GaugesRepository struct {
	data map[domain.GaugeName]domain.GaugeValue
}

func (repository *GaugesRepository) Set(ctx context.Context, key domain.GaugeName, value domain.GaugeValue) error {
	repository.data[key] = value
	return nil
}

func (repository *GaugesRepository) Get(ctx context.Context, key domain.GaugeName) (domain.Gauge, error) {
	value := repository.data[key]
	return domain.Gauge{Name: key, Value: value}, nil
}

func (repository *GaugesRepository) Exists(ctx context.Context, key domain.GaugeName) bool {
	_, found := repository.data[key]
	return found
}

func (repository *GaugesRepository) GetAll(ctx context.Context) ([]domain.Gauge, error) {
	result := make([]domain.Gauge, 0)
	for name, value := range repository.data {
		result = append(result, domain.Gauge{Name: name, Value: value})
	}
	return result, nil
}

var _ services.GaugesRepository = (*GaugesRepository)(nil)
