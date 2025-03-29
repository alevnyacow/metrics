package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Takes data-layer as input and returns
// handler for obtaining all metric values.
func handleGetAllMetrics(metricsRepository datalayer.MetricsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allMetricsJSON, marshalingError := json.Marshal(metricsRepository.AllMetrics())
		if marshalingError != nil {
			marshalingErrorResponse(marshalingError, w)
			return
		}
		w.Write(allMetricsJSON)
	}
}

// Response in case of having error while
// building JSON response data.
func marshalingErrorResponse(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
