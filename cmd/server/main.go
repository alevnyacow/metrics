package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/memstorage"
)

func metricsServeMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	dl := memstorage.NewMemStorage()
	metricsAPIInjector := api.NewMetricsAPIInjector(dl)
	metricsAPIInjector(mux)
	return
}

func apiAddress() string {
	return "localhost:8080"
}

func main() {
	http.ListenAndServe(apiAddress(), metricsServeMux())
}
