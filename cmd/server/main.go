package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/infrastructure/memstorage"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/go-chi/chi/v5"
)

func main() {
	configs := config.ParseServerConfigs()

	inMemoryCountersRepository := memstorage.NewCountersRepository()
	inMemoryGaugesRepository := memstorage.NewGaugesRepository()

	countersService := services.NewCountersService(inMemoryCountersRepository)
	gaugesService := services.NewGaugesService(inMemoryGaugesRepository)

	chiRouter := chi.NewRouter()
	apiController := api.NewController(countersService, gaugesService)
	apiController.AddInChiMux(chiRouter)

	http.ListenAndServe(configs.APIHost, chiRouter)
}
