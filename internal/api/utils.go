package api

import (
	"fmt"
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

const UpdateLinkRoot = "update"
const ValueLinkRoot = "value"

const GaugeLinkPath = "gauge"
const CounterLinkPath = "counter"

const typePathParam = "type"
const namePathParam = "name"
const valuePathParam = "value"

// Generate routes for metric API controller.
func routes() (update string, getMetric string, getAllMetrics string) {
	update = fmt.Sprintf("/%s/{%s}/{%s}/{%s}", UpdateLinkRoot, typePathParam, namePathParam, valuePathParam)
	getMetric = fmt.Sprintf("/%s/{%s}/{%s}", ValueLinkRoot, typePathParam, namePathParam)
	getAllMetrics = "/"
	return
}

// Parse data from "type" path parameter and return
// definite metric type or nothing if we could not map
// this result from path parameter value.
func parseMetricType(request *http.Request) (metricType datalayer.MetricType, success bool) {
	pathParamToMetricType := map[string]datalayer.MetricType{
		GaugeLinkPath:   datalayer.GaugeMetricType,
		CounterLinkPath: datalayer.CounterMetricType,
	}
	metricTypeFromPath := request.PathValue(typePathParam)
	metricType, success = pathParamToMetricType[metricTypeFromPath]
	return
}

// Response in case of metric type was not provided by
// client or provided metric type is unknown for server.
func unknownMetricTypeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

// Response in case of client provided existing
// metric type but there is not metric of requested
// type with requested name.
func nonExistingMetricOfKnownTypeResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
