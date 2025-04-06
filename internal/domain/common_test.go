package domain_test

import (
	"testing"

	"github.com/alevnyacow/metrics/internal/domain"
)

func TestMetricModelFromCounter(t *testing.T) {
	counterName := domain.CounterName("test_counter")
	counterValue := domain.CounterValue(10)
	counter := domain.Counter{Name: counterName, Value: counterValue}
	metricFromCounter := counter.ToMetricModel()
	if metricFromCounter.Name != string(counterName) {
		t.Errorf("Wrong metric name - expected %v, got %v", counterName, metricFromCounter.Name)
	}
	counterValueAsString := counterValue.ToString()
	if metricFromCounter.Value != counterValueAsString {
		t.Errorf("Wrong metric value - expected %v, got %v", counterValueAsString, metricFromCounter.Value)
	}
	if metricFromCounter.Type != domain.CounterMetricType {
		t.Errorf("Wrong metric type - expected %v, got %v", domain.CounterMetricType, metricFromCounter.Type)
	}
}

func TestMetricModelFromGauge(t *testing.T) {
	gaugeName := domain.GaugeName("test_counter")
	gaugeValue := domain.GaugeValue(10)
	gauge := domain.Gauge{Name: gaugeName, Value: gaugeValue}
	metricFromCounter := gauge.ToMetricModel()
	if metricFromCounter.Name != string(gaugeName) {
		t.Errorf("Wrong metric name - expected %v, got %v", gaugeName, metricFromCounter.Name)
	}
	counterValueAsString := gaugeValue.ToString()
	if metricFromCounter.Value != counterValueAsString {
		t.Errorf("Wrong metric value - expected %v, got %v", counterValueAsString, metricFromCounter.Value)
	}
	if metricFromCounter.Type != domain.GaugeMetricType {
		t.Errorf("Wrong metric type - expected %v, got %v", domain.CounterMetricType, metricFromCounter.Type)
	}
}
