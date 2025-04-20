package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleGetMetricValueByPathValue(w http.ResponseWriter, r *http.Request) {
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
		counter, counterWasFound := controller.countersService.GetByKey(domain.CounterName(metricName))
		if !counterWasFound {
			nonExistingMetricOfKnownTypeResponse()(w, r)
			return
		}
		w.Write([]byte(counter.Value))
	case domain.GaugeMetricType:
		gauge, gaugeWasFound := controller.gaugesService.GetByKey(domain.GaugeName(metricName))
		if !gaugeWasFound {
			nonExistingMetricOfKnownTypeResponse()(w, r)
			return
		}
		w.Write([]byte(gauge.Value))
	default:
		unknownMetricTypeResponse()(w, r)
	}
}
