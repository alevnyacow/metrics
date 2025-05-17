package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/retries"
)

func (controller *MetricsController) handleUpdateMultipleMetrics(w http.ResponseWriter, r *http.Request) {
	controller.mutex.Lock()
	defer controller.mutex.Unlock()
	decoder := json.NewDecoder(r.Body)
	var payload []Metric
	err := decoder.Decode(&payload)
	if err != nil {
		marshalingErrorResponse(err)(w, r)
		return
	}
	metrics := make([]domain.Metric, 0)
	for _, item := range payload {
		metrics = append(metrics, item.toDomain())
	}

	error := retries.WithRetries(func() error { return controller.commonMetricsService.UpdateMetrics(r.Context(), metrics) })
	if error != nil {
		failedUpdatesResponse(error)(w, r)
		return
	}
}
