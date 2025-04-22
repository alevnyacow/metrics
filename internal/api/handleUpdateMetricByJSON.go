package api

import (
	"encoding/json"
	"net/http"

	"github.com/alevnyacow/metrics/internal/domain"
)

func (controller *MetricsController) handleUpdateMetricByJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	payload := Metric{}
	err := decoder.Decode(&payload)
	if err != nil {
		marshalingErrorResponse(err)(w, r)
		return
	}
	if payload.ID == "" {
		nonExistingMetricOfKnownTypeResponse()(w, r)
		return
	}
	switch payload.MType {
	case "gauge":
		value, parsed := domain.GaugeRawFloatValue(*payload.Value).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse()(w, r)
			return
		}
		controller.gaugesService.Set(domain.GaugeName(payload.ID), value)
	case "counter":
		value, parsed := domain.CounterRawIntValue(*payload.Delta).ToValue()
		if !parsed {
			providedIncorrectUpdateValueResponse()(w, r)
			return
		}
		controller.countersService.Update(domain.CounterName(payload.ID), value)
	default:
		unknownMetricTypeResponse()(w, r)
		return
	}

	controller.handleGetMetricByJSON(w, r)
}
