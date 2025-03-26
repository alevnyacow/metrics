package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/memstorage"
	"github.com/go-chi/chi/v5"
)

func main() {
	apiHost := config.ForServer()
	chiRouter := chi.NewRouter()
	datalayer := memstorage.NewMemStorage()
	api.ConfigureChiRouter(chiRouter, datalayer)
	http.ListenAndServe(apiHost, chiRouter)
}
