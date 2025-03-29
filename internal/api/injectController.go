package api

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/go-chi/chi/v5"
)

// Takes pointer to chi mux and data-layer object as
// input, injects metrics API controller in provided
// chi mux instanse.
func InjectMetricControllerInChi(chi *chi.Mux, dl datalayer.DataLayer) {
	update, getMetric, getAllMetrics := routes()
	chi.Post(update, handleUpdateMetric(dl))
	chi.Get(getMetric, handleGetMetric(dl))
	chi.Get(getAllMetrics, handleGetAllMetrics(dl))
}
