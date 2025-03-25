package api

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/go-chi/chi/v5"
)

func ConfigureChiRouter(chi *chi.Mux, dl datalayer.DataLayer) {
	chi.Post("/update/{type}/{name}/{value}", newUpdateHandler(dl))
	chi.Get("/value/{type}/{name}", newMetricValueHandler(dl))
	chi.Get("/", newAllValuesHandler(dl))
}
