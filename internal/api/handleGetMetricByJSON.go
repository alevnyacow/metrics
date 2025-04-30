package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleGetMetricByJSON(w http.ResponseWriter, r *http.Request) {
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
		updatedGauge, exists := controller.gaugesService.GetByKey(domain.GaugeName(payload.ID))
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
		updatedCounter, exists := controller.countersService.GetByKey(domain.CounterName(payload.ID))
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
