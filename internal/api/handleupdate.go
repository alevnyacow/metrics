package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func newUpdateHandler(dl datalayer.DataLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parsing metrics type from path, send 400 status
		// code if failed.
		metricType, metricTypeParsingSuccess := parseMetricType(r)
		if !metricTypeParsingSuccess {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Parsing metrics name from path, send 404 status
		// code if failed.
		metricName, metricNameParsingSuccess := parseMetricName(r)
		if !metricNameParsingSuccess {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Parsing metrics value as string from path, sends 400
		// status code if failed.
		stringValue, stringValueParsingSuccess := parseStringValue(r)
		if !stringValueParsingSuccess {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Counter metric updating logic.
		if metricType == counterMetricType {
			// Parsing counter value from its raw string representation,
			// sends 400 status code if failed.
			metricValue, metricValueParsingSuccess := parseCounterValue(stringValue)
			if !metricValueParsingSuccess {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			// Updating counter metric in data-layer by name and value.
			dl.AddCounterMetric(datalayer.CounterName(metricName), metricValue)
		}
		// Gauge metric updating logic.
		if metricType == gaugeMetricType {
			// Parsing gauge value from its raw string representation,
			// sends 400 status code if failed.
			gaugeValue, gaugeValueParsingSuccess := parseGaugeValue(stringValue)
			if !gaugeValueParsingSuccess {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			// Set new gauge value in data-layer by name and value.
			dl.SetGaugeMetric(datalayer.GaugeName(metricName), gaugeValue)
		}
	}
}
