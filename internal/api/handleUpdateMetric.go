package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Takes data-layer as input and returns handler
// for upserting metric by its type, name and value.
func handleUpdateMetric(dl datalayer.DataLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType, metricTypeParsingSuccess := parseMetricType(r)
		if !metricTypeParsingSuccess {
			unknownMetricTypeResponse(w)
			return
		}
		metricName := r.PathValue(namePathParam)
		if metricName == "" {
			nonExistingMetricOfKnownTypeResponse(w)
			return
		}
		stringValue := r.PathValue(valuePathParam)
		if stringValue == "" {
			notProvidedUpdateValueResponse(w)
			return
		}
		switch metricType {
		case datalayer.CounterMetricType:
			updateCounterMetric(dl, datalayer.CounterName(metricName), stringValue, w)
		case datalayer.GaugeMetricType:
			updateGaugeMetric(dl, datalayer.GaugeName(metricName), stringValue, w)
		default:
			unknownMetricTypeResponse(w)
		}
	}
}

// Tries to update requested counter metric
// by metric name and raw string value.
func updateCounterMetric(
	countersRepository datalayer.CountersRepository,
	counterName datalayer.CounterName,
	stringValue string,
	w http.ResponseWriter,
) {
	counterValue, counterValueParsingSuccess := datalayer.CounterValueFromString(stringValue)
	if !counterValueParsingSuccess {
		providedIncorrectUpdateValueResponse(w)
		return
	}
	countersRepository.AddCounterMetric(counterName, counterValue)
}

// Tries to update requested gauge metric
// by metric name and raw string value.
func updateGaugeMetric(
	gaugesRepository datalayer.GaugesRepository,
	gaugeName datalayer.GaugeName,
	stringValue string,
	w http.ResponseWriter,
) {
	gaugeValue, gaugeValueParsingSuccess := datalayer.GaugeValueFromString(stringValue)
	if !gaugeValueParsingSuccess {
		providedIncorrectUpdateValueResponse(w)
		return
	}
	gaugesRepository.SetGaugeMetric(gaugeName, gaugeValue)
}

// Response in case of user did not provide any
// metric update value.
func notProvidedUpdateValueResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

// Response in case of user provided raw metric update
// value but server could not map this string to
// CounterValue or GaugeValue.
func providedIncorrectUpdateValueResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
