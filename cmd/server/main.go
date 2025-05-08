package main

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/alevnyacow/metrics/internal/store/filestorage"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var countersService *services.CountersService
var gaugesService *services.GaugesService
var healthcheckService *services.HealthcheckService
var configs = config.ParseServerConfigs()
var db *sql.DB

func init() {
	wg := sync.WaitGroup{}
	afterUpdate := func() {}
	fileStorage := filestorage.New(configs.FileStoragePath)
	inMemoryCountersRepository := memstorage.NewCountersRepository()
	inMemoryGaugesRepository := memstorage.NewGaugesRepository()

	saveAllMetricsToFile := func() {
		allMetrics := make([]domain.Metric, 0)
		for _, counter := range inMemoryCountersRepository.GetAll() {
			allMetrics = append(allMetrics, counter.ToMetricModel())
		}
		for _, gauge := range inMemoryGaugesRepository.GetAll() {
			allMetrics = append(allMetrics, gauge.ToMetricModel())
		}
		fileStorage.Save(allMetrics)
	}

	if configs.Restore {
		data, err := fileStorage.Load()
		if err == nil {
			for _, metric := range data {
				if metric.IsCounter() {
					value, parsed := domain.CounterRawValue(metric.Value).ToValue()
					if parsed && metric.Name != "" {
						inMemoryCountersRepository.Set(domain.CounterName(metric.Name), value)
					}
				}
				if metric.IsGauge() {
					value, parsed := domain.GaugeRawValue(metric.Value).ToValue()
					if parsed && metric.Name != "" {
						inMemoryGaugesRepository.Set(domain.GaugeName(metric.Name), value)
					}
				}
			}
		}

		if configs.StoreInterval == 0 {
			afterUpdate = saveAllMetricsToFile
		}

		if configs.StoreInterval > 0 {
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()
				for {
					time.Sleep(time.Duration(configs.StoreInterval) * time.Second)
					saveAllMetricsToFile()
				}
			}()
		}
	}

	database, err := sql.Open("postgres", configs.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}

	db = database
	countersService = services.NewCountersService(inMemoryCountersRepository, afterUpdate)
	gaugesService = services.NewGaugesService(inMemoryGaugesRepository, afterUpdate)
	healthcheckService = services.NewHealtheckService(db)
}

func main() {
	defer func() {
		db.Close()
	}()

	chiRouter := chi.NewRouter()
	apiController := api.NewController(countersService, gaugesService, healthcheckService)
	apiController.AddInChiMux(chiRouter)
	server := &http.Server{
		Addr:    configs.APIHost,
		Handler: chiRouter,
	}
	serverStartingError := server.ListenAndServe()
	if serverStartingError != nil {
		log.Err(serverStartingError).Msg("Could not start metrics server")
	}
}
