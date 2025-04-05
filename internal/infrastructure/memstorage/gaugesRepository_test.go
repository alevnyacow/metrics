package memstorage

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
)

var gaugeName = domain.GaugeName("test_gauge")
var gaugeValue = domain.GaugeValue(25.25)

func TestGaugeExists(t *testing.T) {
	gaugesRepository := NewGaugesRepository()
	gaugesRepository.Set(gaugeName, gaugeValue)
	exists := gaugesRepository.Exists(gaugeName)
	if !exists {
		t.Error("Cound not find existing counter")
	}
}

func TestGaugeDoesNotExist(t *testing.T) {
	gaugesRepository := NewGaugesRepository()
	exists := gaugesRepository.Exists(gaugeName)
	if exists {
		t.Error("Found non existing counter")
	}
}

func TestGaugeValue(t *testing.T) {
	gaugesRepository := NewGaugesRepository()
	gaugesRepository.Set(gaugeName, gaugeValue)
	gauge := gaugesRepository.Get(gaugeName)
	if gauge.Value != gaugeValue {
		t.Errorf("Expected %f, got %f", gauge.Value, gaugeValue)
	}
}

func TestGaugesAreEmptyAfterCreation(t *testing.T) {
	gaugesRepository := NewGaugesRepository()
	gauges := gaugesRepository.GetAll()
	if len(gauges) != 0 {
		t.Error("Not empty after creation")
	}
}
