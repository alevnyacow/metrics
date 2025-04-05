package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) updateMetric(w http.ResponseWriter, r *http.Request) {
	metricType, metricTypeParsingSuccess := parseMetricType(r)
	if !metricTypeParsingSuccess {
		unknownMetricTypeResponse()(w, r)
		return
	}
	metricName := r.PathValue(namePathParam)
	if metricName == "" {
		nonExistingMetricOfKnownTypeResponse()(w, r)
		return
	}
	stringValue := r.PathValue(valuePathParam)
	if stringValue == "" {
		notProvidedUpdateValueResponse()(w, r)
		return
	}
	switch metricType {
	case domain.CounterMetricType:
		controller.updateCounter(domain.CounterName(metricName), domain.CounterRawValue(stringValue))(w, r)
	case domain.GaugeMetricType:
		controller.updateGauge(domain.GaugeName(metricName), domain.GaugeRawValue(stringValue))(w, r)
	default:
		unknownMetricTypeResponse()(w, r)
	}
}

func (controller *MetricsController) updateCounter(name domain.CounterName, rawValue domain.CounterRawValue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		success := controller.countersService.SetWithRawValue(name, rawValue)
		if !success {
			providedIncorrectUpdateValueResponse()(w, r)
		}
	}
}

func (controller *MetricsController) updateGauge(name domain.GaugeName, rawValue domain.GaugeRawValue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		success := controller.gaugesService.SetWithRawValue(name, rawValue)
		if !success {
			providedIncorrectUpdateValueResponse()(w, r)
		}
	}
}
