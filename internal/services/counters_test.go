package services_test

import (
	"context"
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
)

var counterName = domain.CounterName("test_counter")
var counterValue = domain.CounterValue(100)
var ctx = context.Background()

func TestCounterValue(t *testing.T) {
	countersRepo := memstorage.NewCountersRepository()
	countersService := services.NewCountersService(countersRepo, func() {})
	countersService.Update(ctx, counterName, counterValue)
	counter, found, err := countersService.GetByKey(ctx, counterName)
	if err != nil {
		t.Error("Error where should not")
	}
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
