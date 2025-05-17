package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleGetMetricByJSON(w http.ResponseWriter, r *http.Request) {
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
	case config.GaugeType:
		updatedGauge, exists, err := controller.gaugesService.GetByKey(r.Context(), domain.GaugeName(payload.ID))
		if err != nil {
			serviceErrorResponse(err)(w, r)
		}
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

	case config.CounterType:
		updatedCounter, exists, err := controller.countersService.GetByKey(r.Context(), domain.CounterName(payload.ID))
		if err != nil {
			serviceErrorResponse(err)(w, r)
		}
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
