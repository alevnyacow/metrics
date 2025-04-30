package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

// parseMetricType tries to parse metric type from path
// values. If there was no provided value or it does not
// correspond to known metric types, success is returned as false.
func parseMetricType(request *http.Request) (metricType domain.MetricType, rawValue string, success bool) {
	pathParamToMetricType := map[string]domain.MetricType{
		GaugeLinkPath:   domain.GaugeMetricType,
		CounterLinkPath: domain.CounterMetricType,
	}
	rawValue = request.PathValue(typePathParam)
	metricType, success = pathParamToMetricType[rawValue]
	return
}
