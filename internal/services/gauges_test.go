package services_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
)

var gaugeName = domain.GaugeName("test_counter")
var gaugeValue = domain.GaugeValue(100)

func TestGaugeValue(t *testing.T) {
	gaugesRepo := memstorage.NewGaugesRepository()
	gaugesService := services.NewGaugesService(gaugesRepo, func() {})
	gaugesService.Set(ctx, gaugeName, gaugeValue)
	gauge, found := gaugesService.GetByKey(ctx, gaugeName)
	if !found {
		t.Error("Have not found existing counter")
	}
	if gauge.Name != string(gaugeName) {
		t.Errorf("Wrong name - expected %s, got %s", gaugeName, gauge.Name)
	}
	if gauge.Value != gaugeValue.ToString() {
		t.Errorf("Wrong gauge value - expected %v, got %v", gaugeValue, gauge.Value)
	}
}
