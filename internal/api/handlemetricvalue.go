package api

import (
	"net/http"
	"strconv"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func newMetricValueHandler(dl datalayer.DataLayer) http.HandlerFunc {
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
		metricName, _ := parseMetricName(request)
		if metricName == "" {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}
		// Counter metric obtaining logic.
		if metricType == counterMetricType {
			metricValue, wasFound := dl.GetCounterValue(datalayer.CounterName(metricName))
			// If requested counter was not found, return 404 status code.
			if !wasFound {
				responseWriter.WriteHeader(http.StatusNotFound)
				return
			}
			// Response with requested counter value.
			responseWriter.Write([]byte(strconv.FormatInt(int64(metricValue), 10)))
			return
		}
		// Gauge metric obtaining logic.
		if metricType == gaugeMetricType {
			metricValue, wasFound := dl.GetGaugeValue(datalayer.GaugeName(metricName))
			// If requested gauge was not found, return 404 status code.
			if !wasFound {
				responseWriter.WriteHeader(http.StatusNotFound)
				return
			}
			// Response with requested gauge value.
			responseWriter.Write([]byte(strconv.FormatFloat(float64(metricValue), 'f', -1, 64)))
		}
	}
}
