// Package api is responsible for metrics WEB API.
package api

import (
	"sync"

	metricsMiddleware "github.com/alevnyacow/metrics/internal/middleware"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// MetricsController provides functionality of metrics
// updating and obtaining metric values (for one metric
// of any type or for all metrics) via WEB API. Compatible
// with Chi.
type MetricsController struct {
	countersService      *services.CountersService
	gaugesService        *services.GaugesService
	healthcheckService   *services.HealthcheckService
	commonMetricsService services.CommonMetricsService
	mutex                *sync.RWMutex
}

func NewController(
	countersService *services.CountersService,
	gaugesService *services.GaugesService,
	healthcheckService *services.HealthcheckService,
	commonMetricsService services.CommonMetricsService,
	mutex *sync.RWMutex,
) *MetricsController {
	return &MetricsController{
		countersService:      countersService,
		gaugesService:        gaugesService,
		healthcheckService:   healthcheckService,
		commonMetricsService: commonMetricsService,
		mutex:                mutex,
	}
}

// AddInChiMux takes Chi mux by pointer as a parameter and
// after this function is called, provided Chi mux has all
// handlers from API Metrics Controller.
func (controller *MetricsController) AddInChiMux(chi *chi.Mux) {
	update, updateWithJSON, getMetric, getAllMetrics, getByJSON, ping, updates := routes()
	chi.Use(middleware.Compress(5, "text/html", "application/json"))
	chi.Use(metricsMiddleware.WithLogging)
	chi.Use(metricsMiddleware.WithGzip)
	chi.Get(getMetric, controller.handleGetMetricValueByPathValue)
	chi.Get(getAllMetrics, controller.handleGetAllMetrics)
	chi.Post(update, controller.updateMetricByPathValues)
	chi.Post(updateWithJSON, controller.handleUpdateMetricByJSON)
	chi.Post(getByJSON, controller.handleGetMetricByJSON)
	chi.Post(updates, controller.handleUpdateMultipleMetrics)
	chi.Get(ping, controller.handlePing)
}
