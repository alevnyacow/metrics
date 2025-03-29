package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/memstorage"
	"github.com/go-chi/chi/v5"
)

func main() {
	configs := config.ParseServerConfigs()
	chiRouter := chi.NewRouter()
	datalayer := memstorage.NewMemStorage()
	api.InjectMetricControllerInChi(chiRouter, datalayer)
	http.ListenAndServe(configs.APIHost, chiRouter)
}
