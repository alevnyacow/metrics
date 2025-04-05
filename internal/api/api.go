// Package api is responsible for metrics web api.
//
// It exports MetricsController which can be created
// via "NewController" constructor. This controller
// is compatible with Chi.
package api

import (
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/go-chi/chi/v5"
)

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
	update, getMetric, getAllMetrics := routes()
	chi.Post(update, controller.updateMetric)
	chi.Get(getMetric, controller.getMetric)
	chi.Get(getAllMetrics, controller.getAllMetrics)
}
