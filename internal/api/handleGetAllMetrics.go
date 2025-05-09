package api

import (
	"encoding/json"
	"net/http"
)

func (controller *MetricsController) handleGetAllMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := append(
		controller.countersService.GetAll(r.Context()),
		controller.gaugesService.GetAll(r.Context())...,
	)
	allMetricsJSON, marshalingError := json.Marshal(metrics)
	if marshalingError != nil {
		marshalingErrorResponse(marshalingError)(w, r)
		return
	}
	w.Write(allMetricsJSON)
}
