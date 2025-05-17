package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleGetMetricValueByPathValue(w http.ResponseWriter, r *http.Request) {
	metricType, rawValue, foundMetricType := parseMetricType(r)
	if !foundMetricType {
		unknownMetricTypeResponse(rawValue)(w, r)
		return
	}
	metricName := r.PathValue(namePathParam)
	if metricName == "" {
		nonExistingMetricOfKnownTypeResponse(metricName)(w, r)
		return
	}
	switch metricType {
	case domain.CounterMetricType:
		counter, counterWasFound, err := controller.countersService.GetByKey(r.Context(), domain.CounterName(metricName))
		if err != nil {
			serviceErrorResponse(err)(w, r)
			return
		}
		if !counterWasFound {
			nonExistingMetricOfKnownTypeResponse(metricName)(w, r)
			return
		}
		w.Write([]byte(counter.Value))
	case domain.GaugeMetricType:
		gauge, gaugeWasFound, err := controller.gaugesService.GetByKey(r.Context(), domain.GaugeName(metricName))
		if err != nil {
			serviceErrorResponse(err)(w, r)
			return
		}
		if !gaugeWasFound {
			nonExistingMetricOfKnownTypeResponse(metricName)(w, r)
			return
		}
		w.Write([]byte(gauge.Value))
	default:
		unknownMetricTypeResponse(rawValue)(w, r)
	}
}
