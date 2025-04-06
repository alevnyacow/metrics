package api

import (
	"encoding/json"
	"net/http"
)

func (controller *MetricsController) getAllMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := append(
		controller.countersService.GetAll(),
		controller.gaugesService.GetAll()...,
	)
	allMetricsJSON, marshalingError := json.Marshal(metrics)
	if marshalingError != nil {
		marshalingErrorResponse(marshalingError)(w, r)
		return
	}
	w.Write(allMetricsJSON)
}
