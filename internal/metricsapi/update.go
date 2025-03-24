package metricsapi

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func newUpdateMetricsDataHandler(dl datalayer.MetricsDataLayer) http.Handler {
	// POST "/update/{type}/{name}/{value}"
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}

		metricType, foundMetricType := getMetricTypeFromRequest(request)

		if !foundMetricType {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}

		if metricType == datalayer.COUNTER_METRIC {
			counterPayload, counterPayloadParsingSuccess := getCounterPayloadFromRequest(request)
			if !counterPayloadParsingSuccess {
				responseWriter.WriteHeader(http.StatusBadRequest)
				return
			}
			dl.AddCounterMetric(counterPayload.name, counterPayload.valueCounter)
		}

		gaugePayload, gaugePayloadParsingSuccess := getGaugePayloadFromRequest(request)

		if !gaugePayloadParsingSuccess {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}

		dl.SetGaugeMetric(gaugePayload.name, gaugePayload.valueCounter)
		responseWriter.WriteHeader(http.StatusOK)
	})
}
