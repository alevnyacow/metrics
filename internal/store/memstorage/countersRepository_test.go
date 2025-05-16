package memstorage_test

import (
	"context"
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
)

var counterName = domain.CounterName("test_counter")
var counterValue = domain.CounterValue(25)
var ctx = context.Background()

func TestCounterExists(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	countersRepository.Set(ctx, counterName, counterValue)
	exists, err := countersRepository.Exists(ctx, counterName)
	if err != nil {
		t.Error("Error where should not")
	}
	if !exists {
		t.Error("Cound not find existing counter")
	}
}

func TestCounterDoesNotExist(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	exists, err := countersRepository.Exists(ctx, counterName)
	if err != nil {
		t.Error("Error where should not")
	}
	if exists {
		t.Error("Found non existing counter")
	}
}

func TestCounterValue(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	countersRepository.Set(ctx, counterName, counterValue)
	foundCounterValue, err := countersRepository.GetValue(ctx, counterName)
	if err != nil {
		t.Error("Error where should not")
	}
	if foundCounterValue != counterValue {
		t.Errorf("Expected %d, got %d", counterValue, foundCounterValue)
	}
}

func TestCountersAreEmptyAfterCreation(t *testing.T) {
	countersRepository := memstorage.NewCountersRepository()
	counters, err := countersRepository.GetAll(ctx)
	if err != nil {
		t.Error("Error where should not")
	}
	if len(counters) != 0 {
		t.Error("Memstorage counters are not empty after creation")
	}
}
