package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) updateMetricByPathValues(w http.ResponseWriter, r *http.Request) {
	controller.mutex.Lock()
	defer controller.mutex.Unlock()

	metricType, rawValue, metricTypeParsingSuccess := parseMetricType(r)
	if !metricTypeParsingSuccess {
		unknownMetricTypeResponse(rawValue)(w, r)
		return
	}
	metricName := r.PathValue(namePathParam)
	if metricName == "" {
		nonExistingMetricOfKnownTypeResponse(metricName)(w, r)
		return
	}
	stringValue := r.PathValue(valuePathParam)
	if stringValue == "" {
		notProvidedUpdateValueResponse()(w, r)
		return
	}
	switch metricType {
	case domain.CounterMetricType:
		value, parsed := domain.CounterRawValue(stringValue).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse(stringValue)(w, r)
		}
		controller.countersService.Update(r.Context(), domain.CounterName(metricName), value)
	case domain.GaugeMetricType:
		value, parsed := domain.GaugeRawValue(stringValue).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse(stringValue)(w, r)
		}
		controller.gaugesService.Set(r.Context(), domain.GaugeName(metricName), value)
	default:
		unknownMetricTypeResponse(rawValue)(w, r)
	}
}
