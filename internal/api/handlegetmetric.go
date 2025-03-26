package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Takes data-layer as input and returns handler
// for obtaining one metric by its type and name.
func handleGetMetric(dl datalayer.DataLayer) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		// Parsing metrics type from path, send 400 status
		// code if failed.
		metricType, foundMetricType := parseMetricType(request)
		if !foundMetricType {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}
		// Parsing metrics name from path, send 404 status
		// code if failed.
		metricName := request.PathValue(namePathParam)
		if metricName == "" {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}
		// Counter metric obtaining logic.
		if metricType == counterMetricType {
			counterValue, counterWasFound := dl.GetCounterValue(datalayer.CounterName(metricName))
			// If requested counter was not found, return 404 status code.
			if !counterWasFound {
				responseWriter.WriteHeader(http.StatusNotFound)
				return
			}
			// Response with requested counter value.
			responseWriter.Write([]byte(datalayer.CounterValueToString(counterValue)))
		}
		// Gauge metric obtaining logic.
		if metricType == gaugeMetricType {
			gaugeValue, gaugeWasFound := dl.GetGaugeValue(datalayer.GaugeName(metricName))
			// If requested gauge was not found, return 404 status code.
			if !gaugeWasFound {
				responseWriter.WriteHeader(http.StatusNotFound)
				return
			}
			// Response with requested gauge value.
			responseWriter.Write([]byte(datalayer.GaugeValueToString(gaugeValue)))
		}
	}
}
