package memstorage

import (
	"context"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
)

// CounterRepository is in-memory implementation
// of CounterRepository interface.
type CountersRepository struct {
	data map[domain.CounterName]domain.CounterValue
}

func (repository *CountersRepository) Set(ctx context.Context, key domain.CounterName, value domain.CounterValue) error {
	repository.data[key] = value
	return nil
}

func (repository *CountersRepository) Get(ctx context.Context, key domain.CounterName) (domain.Counter, error) {
	value := repository.data[key]
	return domain.Counter{Name: key, Value: value}, nil
}

func (repository *CountersRepository) GetValue(ctx context.Context, key domain.CounterName) (domain.CounterValue, error) {
	return repository.data[key], nil
}

func (repository *CountersRepository) Exists(ctx context.Context, key domain.CounterName) (bool, error) {
	_, found := repository.data[key]
	return found, nil
}

func (repository *CountersRepository) GetAll(ctx context.Context) ([]domain.Counter, error) {
	result := make([]domain.Counter, 0)
	for name, value := range repository.data {
		result = append(result, domain.Counter{Name: name, Value: value})
	}
	return result, nil
}

var _ services.CountersRepository = (*CountersRepository)(nil)
