// Package api is responsible for metrics WEB API.
package api

import (
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// MetricsController provides functionality of metrics
// updating and obtaining metric values (for one metric
// of any type or for all metrics) via WEB API. Compatible
// with Chi.
type MetricsController struct {
	countersService *services.CountersService
	gaugesService   *services.GaugesService
}

func NewController(
	countersService *services.CountersService,
	gaugesService *services.GaugesService,
) *MetricsController {
	return &MetricsController{
		countersService: countersService,
		gaugesService:   gaugesService,
	}
}

// AddInChiMux takes Chi mux by pointer as a parameter and
// after this function is called, provided Chi mux has all
// handlers from API Metrics Controller.
func (controller *MetricsController) AddInChiMux(chi *chi.Mux) {
	update, updateWithJSON, getMetric, getAllMetrics, getByJSON := routes()
	chi.Use(middleware.Compress(5, "text/html", "application/json"))
	chi.Use(withLogging)
	chi.Get(getMetric, controller.handleGetMetricValueByPathValue)
	chi.With(GzipMiddleware).Get(getAllMetrics, controller.handleGetAllMetrics)
	chi.Post(update, controller.updateMetricByPathValues)
	chi.With(GzipMiddleware).Post(updateWithJSON, controller.handleUpdateMetricByJSON)
	chi.Post(getByJSON, controller.handleGetMetricByJSON)
}
