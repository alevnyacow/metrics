package memstorage_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
)

var gaugeName = domain.GaugeName("test_gauge")
var gaugeValue = domain.GaugeValue(25.25)

func TestGaugeExists(t *testing.T) {
	gaugesRepository := memstorage.NewGaugesRepository()
	gaugesRepository.Set(ctx, gaugeName, gaugeValue)
	exists, err := gaugesRepository.Exists(ctx, gaugeName)
	if err != nil {
		t.Error("Error where should not")
	}
	if !exists {
		t.Error("Cound not find existing gauge")
	}
}

func TestGaugeDoesNotExist(t *testing.T) {
	gaugesRepository := memstorage.NewGaugesRepository()
	exists, err := gaugesRepository.Exists(ctx, gaugeName)
	if err != nil {
		t.Error("Error where should not")
	}
	if exists {
		t.Error("Found non existing gauge")
	}
}

func TestGaugeValue(t *testing.T) {
	gaugesRepository := memstorage.NewGaugesRepository()
	gaugesRepository.Set(ctx, gaugeName, gaugeValue)
	gauge, err := gaugesRepository.Get(ctx, gaugeName)
	if err != nil {
		t.Error("Error where should not")
	}
	if gauge.Value != gaugeValue {
		t.Errorf("Expected %f, got %f", gauge.Value, gaugeValue)
	}
}

func TestGaugesAreEmptyAfterCreation(t *testing.T) {
	gaugesRepository := memstorage.NewGaugesRepository()
	gauges, err := gaugesRepository.GetAll(ctx)
	if err != nil {
		t.Error("Error where should not")
	}
	if len(gauges) != 0 {
		t.Error("Memstorage gauges are not empty after creation")
	}
}
