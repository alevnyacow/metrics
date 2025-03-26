package api

import (
	"testing"
)

func TestRoutePathValues(t *testing.T) {
	update, getMetric, getAllMetrics := routes()

	expectedUpdateValue := "/update/{type}/{name}/{value}"
	if update != expectedUpdateValue {
		t.Errorf("Expected %s, got %s", expectedUpdateValue, update)
	}

	expectedGetMetricValue := "/value/{type}/{name}"
	if getMetric != expectedGetMetricValue {
		t.Errorf("Expected %s, got %s", expectedGetMetricValue, getMetric)
	}

	expectedGetAllMetricsValue := "/"
	if getAllMetrics != expectedGetAllMetricsValue {
		t.Errorf("Expected %s, got %s", expectedGetAllMetricsValue, getAllMetrics)
	}
}
