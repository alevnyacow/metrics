package api

import (
	"fmt"
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func newUpdateMetricsDataHandler(dl datalayer.DataLayer) http.Handler {
	// POST "/update/{type}/{name}/{value}"
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			responseWriter.WriteHeader(http.StatusNotFound)
			return
		}

		metricType, foundMetricType := parseMetricTypeFromRequest(request)

		if !foundMetricType {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}

		if metricType == counterMetricType {
			counterPayload, counterPayloadParsingResult := parseCounterPayloadFromRequest(request)
			isInErrorState := processUpdatePathParsingResult(counterPayloadParsingResult, responseWriter)
			if !isInErrorState {
				dl.AddCounterMetric(counterPayload.name, counterPayload.valueCounter)
				responseWriter.WriteHeader(http.StatusOK)
			}
			fmt.Println(dl.GetCounterValue(datalayer.CounterName(counterPayload.name)))
			return
		}

		gaugePayload, gaugePayloadParsingResult := parseGaugePayloadFromRequest(request)
		isInErrorState := processUpdatePathParsingResult(gaugePayloadParsingResult, responseWriter)

		if !isInErrorState {
			dl.SetGaugeMetric(gaugePayload.name, gaugePayload.valueCounter)
			responseWriter.WriteHeader(http.StatusOK)
			return
		}
	})
}
