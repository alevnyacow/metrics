package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/alevnyacow/metrics/internal/metricsapi"
)

func metricsServeMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	dl := datalayer.NewMemStorage()
	metricsAPIInjector := metricsapi.NewMetricsAPIInjector(dl)
	metricsAPIInjector(mux)
	return
}

func apiAddress() string {
	return "localhost:8080"
}

func main() {
	http.ListenAndServe(apiAddress(), metricsServeMux())
}
