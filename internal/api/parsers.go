package api

import (
	"net/http"
)

// Definite metric type.
type metric string

const (
	// Gauge metric.
	gaugeMetricType metric = "GAUGE"
	// Counter metric.
	counterMetricType metric = "COUNTER"
)

// Parse data from "type" path parameter and return
// definite metric type or nothing if we could not map
// this result from path parameter value.
func parseMetricType(request *http.Request) (metricType metric, success bool) {
	pathParamToMetricType := map[string]metric{
		"gauge":   gaugeMetricType,
		"counter": counterMetricType,
	}
	metricTypeFromPath := request.PathValue(typePathParam)
	metricType, success = pathParamToMetricType[metricTypeFromPath]
	return
}
