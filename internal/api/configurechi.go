package api

import (
	"fmt"

	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/go-chi/chi/v5"
)

const typePathParam = "type"
const namePathParam = "name"
const valuePathParam = "value"

func routes() (update string, getMetric string, getAllMetrics string) {
	update = fmt.Sprintf("/update/{%s}/{%s}/{%s}", typePathParam, namePathParam, valuePathParam)
	getMetric = fmt.Sprintf("/value/{%s}/{%s}", typePathParam, namePathParam)
	getAllMetrics = "/"
	return
}

func ConfigureChiRouter(chi *chi.Mux, dl datalayer.DataLayer) {
	update, getMetric, getAllMetrics := routes()
	chi.Post(update, handleUpdateMetric(dl))
	chi.Get(getMetric, handleGetMetric(dl))
	chi.Get(getAllMetrics, handleGetAllMetrics(dl))
}
