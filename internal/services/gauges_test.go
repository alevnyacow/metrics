package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/infrastructure/memstorage"
	"github.com/alevnyacow/metrics/internal/services"
)

var gaugeName = domain.GaugeName("test_counter")
var gaugeRawValue = domain.GaugeRawValue("100")
var gaugeValue = domain.GaugeValue(100)

func TestGaugeValue(t *testing.T) {
	gaugesRepo := memstorage.NewGaugesRepository()
	gaugesService := services.NewGaugesService(gaugesRepo)
	gaugesService.SetWithRawValue(gaugeName, gaugeRawValue)
	counter, found := gaugesService.GetByKey(gaugeName)
	if !found {
		t.Error("Have not found existing counter")
	}
	if counter.Name != string(counterName) {
		t.Errorf("Wrong name - expected %s, got %s", counterName, counter.Name)
	}
	if counter.Value != string(gaugeRawValue) {
		t.Error("Wrong string value representation - expected %w, got %w", gaugeRawValue, counter.Value)
	}
	counterActualValue, parsed := domain.GaugeRawValue(counter.Value).ToValue()
	if !parsed {
		t.Error("Could not parse counter value")
	}
	if counterActualValue != gaugeValue {
		t.Errorf("Wrong parsed counter values - expected %v, got %v", gaugeValue, counterActualValue)
	}
}
