package memstorage_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/infrastructure/memstorage"
)

var counterName = domain.CounterName("test_counter")
var counterValue = domain.CounterValue(25)

func TestCounterExists(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	countersRepository.Set(counterName, counterValue)
	exists := countersRepository.Exists(counterName)
	if !exists {
		t.Error("Cound not find existing counter")
	}
}

func TestCounterDoesNotExist(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	exists := countersRepository.Exists(counterName)
	if exists {
		t.Error("Found non existing counter")
	}
}

func TestCounterValue(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	countersRepository.Set(counterName, counterValue)
	foundCounterValue := countersRepository.GetValue(counterName)
	if foundCounterValue != counterValue {
		t.Errorf("Expected %d, got %d", counterValue, foundCounterValue)
	}
}

func TestCountersAreEmptyAfterCreation(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	counters := countersRepository.GetAll()
	if len(counters) != 0 {
		t.Error("Memstorage counters are not empty after creation")
	}
}
