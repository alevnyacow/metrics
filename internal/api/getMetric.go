package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) getMetric(w http.ResponseWriter, r *http.Request) {
	metricType, foundMetricType := parseMetricType(r)
	if !foundMetricType {
		unknownMetricTypeResponse()(w, r)
		return
	}
	metricName := r.PathValue(namePathParam)
	if metricName == "" {
		nonExistingMetricOfKnownTypeResponse()(w, r)
		return
	}
	switch metricType {
	case domain.CounterMetricType:
		controller.getCounter(domain.CounterName(metricName))(w, r)
	case domain.GaugeMetricType:
		controller.getGauge(domain.GaugeName(metricName))(w, r)
	default:
		unknownMetricTypeResponse()(w, r)
	}
}

func (controller *MetricsController) getCounter(name domain.CounterName) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter, counterWasFound := controller.countersService.GetByKey(name)
		if !counterWasFound {
			nonExistingMetricOfKnownTypeResponse()(w, r)
			return
		}
		w.Write([]byte(counter.Value))
	}
}

func (controller *MetricsController) getGauge(name domain.GaugeName) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gauge, gaugeWasFound := controller.gaugesService.GetByKey(name)
		if !gaugeWasFound {
			nonExistingMetricOfKnownTypeResponse()(w, r)
			return
		}
		w.Write([]byte(gauge.Value))

	}
}
