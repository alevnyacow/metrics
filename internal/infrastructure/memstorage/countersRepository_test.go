package memstorage

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
)

var counterName = domain.CounterName("test_counter")
var counterValue = domain.CounterValue(25)

func TestCounterExists(t *testing.T) {
	countersRepository := NewCountersRepository()
	countersRepository.Set(counterName, counterValue)
	exists := countersRepository.Exists(counterName)
	if !exists {
		t.Error("Cound not find existing counter")
	}
}

func TestCounterDoesNotExist(t *testing.T) {
	countersRepository := NewCountersRepository()
	exists := countersRepository.Exists(counterName)
	if exists {
		t.Error("Found non existing counter")
	}
}

func TestCounterValue(t *testing.T) {
	countersRepository := NewCountersRepository()
	countersRepository.Set(counterName, counterValue)
	foundCounterValue := countersRepository.GetValue(counterName)
	if foundCounterValue != counterValue {
		t.Errorf("Expected %d, got %d", counterValue, foundCounterValue)
	}
}

func TestCountersAreEmptyAfterCreation(t *testing.T) {
	countersRepository := NewCountersRepository()
	counters := countersRepository.GetAll()
	if len(counters) != 0 {
		t.Error("Not empty after creation")
	}
}
