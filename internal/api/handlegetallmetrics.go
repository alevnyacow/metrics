package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func handleGetAllMetrics(dl datalayer.DataLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get JSON representation of all metrics.
		allMetricsJSON, marshalingError := json.Marshal(dl.AllMetrics())
		// Return 500 status code and error description,
		// if not succeed.
		if marshalingError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(marshalingError.Error()))
			return
		}
		// Response with JSON representation of
		// all existing metrics.
		w.Write(allMetricsJSON)
	}
}
