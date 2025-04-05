package memstorage

import (
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
)

// GaugesRepository is in-memory implementation
// of GaugesRepository interface.
type GaugesRepository struct {
	data map[domain.GaugeName]domain.GaugeValue
}

func (repository *GaugesRepository) Set(key domain.GaugeName, value domain.GaugeValue) {
	repository.data[key] = value
}

func (repository *GaugesRepository) Get(key domain.GaugeName) domain.Gauge {
	value := repository.data[key]
	return domain.Gauge{Name: key, Value: value}
}

func (repository *GaugesRepository) Exists(key domain.GaugeName) bool {
	_, found := repository.data[key]
	return found
}

func (repository *GaugesRepository) GetAll() []domain.Gauge {
	result := make([]domain.Gauge, 0)
	for name, value := range repository.data {
		result = append(result, domain.Gauge{Name: name, Value: value})
	}
	return result
}

var _ services.GaugesRepository = (*GaugesRepository)(nil)
