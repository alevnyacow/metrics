package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleUpdateMetricByJSON(w http.ResponseWriter, r *http.Request) {
	controller.mutex.Lock()
	defer controller.mutex.Unlock()

	decoder := json.NewDecoder(r.Body)
	payload := Metric{}
	err := decoder.Decode(&payload)
	if err != nil {
		marshalingErrorResponse(err)(w, r)
		return
	}
	if payload.ID == "" {
		nonExistingMetricOfKnownTypeResponse(payload.ID)(w, r)
		return
	}
	switch payload.MType {
	case "gauge":
		value, parsed := domain.GaugeRawFloatValue(*payload.Value).ToValue()
		if parsed {
			controller.gaugesService.Set(r.Context(), domain.GaugeName(payload.ID), value)
		}
		updatedGauge, exists := controller.gaugesService.GetByKey(r.Context(), domain.GaugeName(payload.ID))
		if !exists {
			nonExistingMetricOfKnownTypeResponse(payload.ID)(w, r)
			return
		}
		metricDTO := MapDomainMetricToMetricDTO(updatedGauge)
		metricJSON, marshalingError := json.Marshal(metricDTO)
		if marshalingError != nil {
			marshalingErrorResponse(marshalingError)(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(metricJSON)

	case "counter":
		value, parsed := domain.CounterRawIntValue(*payload.Delta).ToValue()
		if parsed {
			controller.countersService.Update(r.Context(), domain.CounterName(payload.ID), value)
		}
		updatedCounter, exists := controller.countersService.GetByKey(r.Context(), domain.CounterName(payload.ID))
		if !exists {
			nonExistingMetricOfKnownTypeResponse(payload.ID)(w, r)
			return
		}
		metricDTO := MapDomainMetricToMetricDTO(updatedCounter)
		metricJSON, marshalingError := json.Marshal(metricDTO)
		if marshalingError != nil {
			marshalingErrorResponse(marshalingError)(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(metricJSON)

	default:
		unknownMetricTypeResponse(payload.MType)(w, r)
	}
}
