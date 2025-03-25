package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func newAllValuesHandler(dl datalayer.DataLayer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allGauges := dl.AllGauges()
		allCounters := dl.AllCounters()
		stringRepresentation := make([]string, 0)
		for _, gaugeMetricDTO := range allGauges {
			value, err := json.Marshal(gaugeMetricDTO)
			if err == nil {
				stringRepresentation = append(stringRepresentation, string(value))
			}
		}
		for _, counterMetricDTO := range allCounters {
			value, err := json.Marshal(counterMetricDTO)
			if err == nil {
				stringRepresentation = append(stringRepresentation, string(value))
			}
		}
		resultBody := strings.Join(stringRepresentation, "\n")
		w.Write([]byte(resultBody))
	}
}
