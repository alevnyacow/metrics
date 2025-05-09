package main

import (
	"context"
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/alevnyacow/metrics/internal/store/dbstorage"
	"github.com/alevnyacow/metrics/internal/store/filestorage"
	"github.com/alevnyacow/metrics/internal/store/memstorage"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var countersService *services.CountersService
var gaugesService *services.GaugesService
var healthcheckService *services.HealthcheckService
var commonMetricsService services.CommonMetricsService
var configs = config.ParseServerConfigs()
var db *sql.DB
var ctx = context.Background()
var mutex = &sync.Mutex{}

func init() {
	wg := sync.WaitGroup{}

	afterUpdate := func() {}
	if configs.DatabaseConnectionString == "" {
		fileStorage := filestorage.New(configs.FileStoragePath)
		inMemoryCountersRepository := memstorage.NewCountersRepository()
		inMemoryGaugesRepository := memstorage.NewGaugesRepository()

		saveAllMetricsToFile := func() {
			allMetrics := make([]domain.Metric, 0)
			for _, counter := range inMemoryCountersRepository.GetAll(ctx) {
				allMetrics = append(allMetrics, counter.ToMetricModel())
			}
			for _, gauge := range inMemoryGaugesRepository.GetAll(ctx) {
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
							inMemoryCountersRepository.Set(ctx, domain.CounterName(metric.Name), value)
						}
					}
					if metric.IsGauge() {
						value, parsed := domain.GaugeRawValue(metric.Value).ToValue()
						if parsed && metric.Name != "" {
							inMemoryGaugesRepository.Set(ctx, domain.GaugeName(metric.Name), value)
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

		countersService = services.NewCountersService(inMemoryCountersRepository, afterUpdate)
		gaugesService = services.NewGaugesService(inMemoryGaugesRepository, afterUpdate)
		commonMetricsService = services.NewInMemoryCommonMetricsService(countersService, gaugesService)
		healthcheckService = services.NewHealtheckService(nil)
		return
	}

	database, err := sql.Open("postgres", configs.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}
	db = database

	dbCountersRepo := dbstorage.NewCountersRepository(db)
	dbCountersRepo.PrepareDB(ctx)
	dbGaugesRepo := dbstorage.NewGaugesRepository(db)
	dbGaugesRepo.PrepareDB(ctx)

	countersService = services.NewCountersService(
		dbCountersRepo,
		afterUpdate,
	)
	gaugesService = services.NewGaugesService(
		dbGaugesRepo,
		afterUpdate,
	)
	healthcheckService = services.NewHealtheckService(db)
	commonMetricsService = services.NewDbCommonMetricsService(db, dbCountersRepo)
}

func main() {
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	chiRouter := chi.NewRouter()
	apiController := api.NewController(countersService, gaugesService, healthcheckService, commonMetricsService, mutex)
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
