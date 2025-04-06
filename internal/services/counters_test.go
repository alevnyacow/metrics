package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/infrastructure/memstorage"
	"github.com/alevnyacow/metrics/internal/services"
)

var counterName = domain.CounterName("test_counter")
var counterRawValue = domain.CounterRawValue("100")
var counterValue = domain.CounterValue(100)

func TestCounterValue(t *testing.T) {
	countersRepo := memstorage.NewCountersRepository()
	countersService := services.NewCountersService(countersRepo)
	countersService.SetWithRawValue(counterName, counterRawValue)
	counter, found := countersService.GetByKey(counterName)
	if !found {
		t.Error("Have not found existing counter")
	}
	if counter.Name != string(counterName) {
		t.Errorf("Wrong name - expected %s, got %s", counterName, counter.Name)
	}
	if counter.Value != string(counterRawValue) {
		t.Error("Wrong string value representation - expected %w, got %w", counterRawValue, counter.Value)
	}
	counterActualValue, parsed := domain.CounterRawValue(counter.Value).ToValue()
	if !parsed {
		t.Error("Could not parse counter value")
	}
	if counterActualValue != counterValue {
		t.Errorf("Wrong parsed counter values - expected %d, got %d", counterValue, counterActualValue)
	}
}
