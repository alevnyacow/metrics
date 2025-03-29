package memstorage

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func TestMetricsRepositoryWithOneCounter(t *testing.T) {
	memStorage := NewMemStorage()
	memStorage.AddCounterMetric("test", 1)
	allMetrics := memStorage.AllMetrics()
	for _, value := range allMetrics {
		if value.Name != "test" {
			t.Errorf("Expected 'test', got %s", value.Name)
		}
		if value.Value != "1" {
			t.Errorf("Expected '1', got %s", value.Value)
		}
		if value.Type != datalayer.CounterMetricType {
			t.Errorf("Expected %s, got %s", datalayer.CounterMetricType, value.Type)
		}
	}
}

func TestMetricsRepositoryWithOneGauge(t *testing.T) {
	memStorage := NewMemStorage()
	memStorage.SetGaugeMetric("test", 1)
	allMetrics := memStorage.AllMetrics()
	for _, value := range allMetrics {
		if value.Name != "test" {
			t.Errorf("Expected 'test', got %s", value.Name)
		}
		if value.Value != "1" {
			t.Errorf("Expected '1', got %s", value.Value)
		}
		if value.Type != datalayer.GaugeMetricType {
			t.Errorf("Expected %s, got %s", datalayer.GaugeMetricType, value.Type)
		}
	}
}
