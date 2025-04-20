package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/infrastructure/memstorage"
	"github.com/alevnyacow/metrics/internal/services"
)

var counterName = domain.CounterName("test_counter")
var counterValue = domain.CounterValue(100)

func TestCounterValue(t *testing.T) {
	countersRepo := memstorage.NewCountersRepository()
	countersService := services.NewCountersService(countersRepo)
	countersService.Update(counterName, counterValue)
	counter, found := countersService.GetByKey(counterName)
	if !found {
		t.Error("Have not found existing counter")
	}
	if counter.Name != string(counterName) {
		t.Errorf("Wrong name - expected %s, got %s", counterName, counter.Name)
	}
	if counter.Value != counterValue.ToString() {
		t.Errorf("Wrong parsed counter values - expected %d, got %s", counterValue, counter.Value)
	}
}
