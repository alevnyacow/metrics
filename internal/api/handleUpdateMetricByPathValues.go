package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) updateMetricByPathValues(w http.ResponseWriter, r *http.Request) {
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
		value, parsed := domain.CounterRawValue(stringValue).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse()(w, r)
		}
		controller.countersService.Update(domain.CounterName(metricName), value)
	case domain.GaugeMetricType:
		value, parsed := domain.GaugeRawValue(stringValue).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse()(w, r)
		}
		controller.gaugesService.Set(domain.GaugeName(metricName), value)
	default:
		unknownMetricTypeResponse()(w, r)
	}
}
