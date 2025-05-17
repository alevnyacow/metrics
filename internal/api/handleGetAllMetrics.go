package api

import (
	"encoding/json"
	"net/http"
)

func (controller *MetricsController) handleGetAllMetrics(w http.ResponseWriter, r *http.Request) {
	controller.mutex.Lock()
	defer controller.mutex.Unlock()

	counters, errOnGettingCounters := controller.countersService.GetAll(r.Context())
	if errOnGettingCounters != nil {
		serviceErrorResponse(errOnGettingCounters)(w, r)
		return
	}
	gauges, errOnGettingGauges := controller.gaugesService.GetAll(r.Context())
	if errOnGettingGauges != nil {
		serviceErrorResponse(errOnGettingGauges)(w, r)
		return
	}

	metrics := append(
		counters,
		gauges...,
	)
	allMetricsJSON, marshalingError := json.Marshal(metrics)
	if marshalingError != nil {
		marshalingErrorResponse(marshalingError)(w, r)
		return
	}
	w.Write(allMetricsJSON)
}
