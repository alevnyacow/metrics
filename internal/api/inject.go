package api

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Takes data-layer as input and returns function which
// injects Metrics API controller when applied to
// given ServeMux.
func NewMetricsAPIInjector(dl datalayer.MetricsDataLayer) func(mux *http.ServeMux) {
	updateHandler := newUpdateMetricsDataHandler(dl)

	return func(mux *http.ServeMux) {
		mux.Handle("/update/{type}/{name}/{value}", updateHandler)
	}
}
