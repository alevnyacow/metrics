package api

import (
	"testing"
)

func TestRoutePathValues(t *testing.T) {
	update, updateWithJSON, getMetric, getAllMetrics, getByJSON, ping := routes()

	expectedUpdateValue := "/update/{type}/{name}/{value}"
	if update != expectedUpdateValue {
		t.Errorf("Expected %s, got %s", expectedUpdateValue, update)
	}

	expectedUpdateWithJSONValue := "/update/"
	if updateWithJSON != expectedUpdateWithJSONValue {
		t.Errorf("Expected %s, got %s", expectedUpdateWithJSONValue, updateWithJSON)
	}

	expectedGetMetricValue := "/value/{type}/{name}"
	if getMetric != expectedGetMetricValue {
		t.Errorf("Expected %s, got %s", expectedGetMetricValue, getMetric)
	}

	expectedGetAllMetricsValue := "/"
	if getAllMetrics != expectedGetAllMetricsValue {
		t.Errorf("Expected %s, got %s", expectedGetAllMetricsValue, getAllMetrics)
	}

	expectedGetByJSONValue := "/value/"
	if getByJSON != expectedGetByJSONValue {
		t.Errorf("Expected %s, got %s", expectedGetByJSONValue, getByJSON)
	}

	expectedPingValue := "/ping"
	if expectedPingValue != ping {
		t.Error()
	}
}
