package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Takes data-layer object as input and returns handler
// for obtaining one metric by its type and name.
func handleGetMetric(dl datalayer.DataLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType, foundMetricType := parseMetricType(r)
		if !foundMetricType {
			unknownMetricTypeResponse(w)
			return
		}
		metricName := r.PathValue(namePathParam)
		if metricName == "" {
			nonExistingMetricOfKnownTypeResponse(w)
			return
		}
		switch metricType {
		case datalayer.CounterMetricType:
			provideCounterMetric(dl, datalayer.CounterName(metricName), w)
		case datalayer.GaugeMetricType:
			obtainGaugeMetric(dl, datalayer.GaugeName(metricName), w)
		default:
			unknownMetricTypeResponse(w)
		}
	}
}

// Tries to build a response with a requested
// counter metric value by its name.
func provideCounterMetric(
	countersRepository datalayer.CountersRepository,
	counterName datalayer.CounterName,
	w http.ResponseWriter,
) {
	counterValue, counterWasFound := countersRepository.GetCounterValue(counterName)
	if !counterWasFound {
		nonExistingMetricOfKnownTypeResponse(w)
		return
	}
	w.Write([]byte(datalayer.CounterValueToString(counterValue)))
}

// Tries to build a response with a requested
// gauge metric value by its name.
func obtainGaugeMetric(
	gaugesRepository datalayer.GaugesRepository,
	gaugeName datalayer.GaugeName,
	w http.ResponseWriter,
) {
	gaugeValue, gaugeWasFound := gaugesRepository.GetGaugeValue(gaugeName)
	if !gaugeWasFound {
		nonExistingMetricOfKnownTypeResponse(w)
		return
	}
	w.Write([]byte(datalayer.GaugeValueToString(gaugeValue)))
}
