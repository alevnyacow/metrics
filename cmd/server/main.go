package main

import (
	"net/http"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/memstorage"
	"github.com/go-chi/chi/v5"
)

func apiAddress() string {
	return "localhost:8080"
}

func main() {
	chiRouter := chi.NewRouter()
	datalayer := memstorage.NewMemStorage()
	api.ConfigureChiRouter(chiRouter, datalayer)
	http.ListenAndServe(apiAddress(), chiRouter)
}
