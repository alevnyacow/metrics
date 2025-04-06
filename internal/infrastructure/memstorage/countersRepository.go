package memstorage

import (
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
)

// CounterRepository is in-memory implementation
// of CounterRepository interface.
type CountersRepository struct {
	data map[domain.CounterName]domain.CounterValue
}

func (repository *CountersRepository) Set(key domain.CounterName, value domain.CounterValue) {
	repository.data[key] = value
}

func (repository *CountersRepository) Get(key domain.CounterName) domain.Counter {
	value := repository.data[key]
	return domain.Counter{Name: key, Value: value}
}

func (repository *CountersRepository) GetValue(key domain.CounterName) domain.CounterValue {
	return repository.data[key]
}

func (repository *CountersRepository) Exists(key domain.CounterName) bool {
	_, found := repository.data[key]
	return found
}

func (repository *CountersRepository) GetAll() []domain.Counter {
	result := make([]domain.Counter, 0)
	for name, value := range repository.data {
		result = append(result, domain.Counter{Name: name, Value: value})
	}
	return result
}

var _ services.CountersRepository = (*CountersRepository)(nil)
