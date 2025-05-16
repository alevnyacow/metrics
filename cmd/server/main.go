package main

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/alevnyacow/metrics/internal/retries"
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
var ctx = context.Background()
var mutex = &sync.Mutex{}
var wg = sync.WaitGroup{}
var afterUpdate = func() {}

func saveAllMetricsToFile(
	fileStorage *filestorage.FileStorage,
	inMemoryCountersRepository *memstorage.CountersRepository,
	inMemoryGaugesRepository *memstorage.GaugesRepository,
) {
	counters, errOnGettingCounters := inMemoryCountersRepository.GetAll(ctx)
	if errOnGettingCounters != nil {
		log.Err(errOnGettingCounters).Msg("Error on getting counters while saving to file")
		return
	}
	gauges, errOnGettingGauges := inMemoryGaugesRepository.GetAll(ctx)
	if errOnGettingGauges != nil {
		log.Err(errOnGettingGauges).Msg("Error on getting gauges while saving to file")
		return
	}
	counterGaugesLen := len(counters) + len(gauges)
	allMetrics := make([]domain.Metric, counterGaugesLen)

	for i, counter := range counters {
		allMetrics[i] = counter.ToMetricModel()
	}
	for i, gauge := range gauges {
		allMetrics[len(counters)+i] = gauge.ToMetricModel()
	}
	retries.WithRetries(func() error { return fileStorage.Save(allMetrics) })
}

func prepareApplicationWithMemStorage() {
	fileStorage := filestorage.New(configs.FileStoragePath)
	inMemoryCountersRepository := memstorage.NewCountersRepository()
	inMemoryGaugesRepository := memstorage.NewGaugesRepository()

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
			afterUpdate = func() { saveAllMetricsToFile(fileStorage, inMemoryCountersRepository, inMemoryGaugesRepository) }
		}

		if configs.StoreInterval > 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					time.Sleep(time.Duration(configs.StoreInterval) * time.Second)
					saveAllMetricsToFile(fileStorage, inMemoryCountersRepository, inMemoryGaugesRepository)
				}
			}()
		}
	}

	countersService = services.NewCountersService(inMemoryCountersRepository, afterUpdate)
	gaugesService = services.NewGaugesService(inMemoryGaugesRepository, afterUpdate)
	commonMetricsService = services.NewInMemoryCommonMetricsService(countersService, gaugesService)
	healthcheckService = services.NewHealtheckService(nil)

}

func prepareApplicationWithDB(connectionString string) func() {
	db, closeDatabase := dbstorage.InitDatabase(connectionString)
	dbCountersRepo := dbstorage.NewCountersRepository(db)
	retries.WithRetries(func() error { return dbCountersRepo.PrepareDB(ctx) })
	dbGaugesRepo := dbstorage.NewGaugesRepository(db)
	retries.WithRetries(func() error { return dbGaugesRepo.PrepareDB(ctx) })

	countersService = services.NewCountersService(
		dbCountersRepo,
		afterUpdate,
	)
	gaugesService = services.NewGaugesService(
		dbGaugesRepo,
		afterUpdate,
	)
	healthcheckService = services.NewHealtheckService(db)
	commonMetricsService = services.NewDBCommonMetricsService(db, dbCountersRepo)
	return closeDatabase
}

func main() {
	if configs.DatabaseConnectionString == "" {
		prepareApplicationWithMemStorage()
	} else {
		closeDatabase := prepareApplicationWithDB(configs.DatabaseConnectionString)
		defer closeDatabase()
	}

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
	wg.Wait()
}
