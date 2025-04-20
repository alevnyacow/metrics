package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/services"
)

type Callback func()

func main() {
	mutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	configs := config.ParseAgentConfigs()
	updateURL := api.MetricUpdateByJSONRoute(configs.APIHost)
	metricsCollectionService := services.NewMetricsCollectionService()
	repetetiveGoroutineCreator := newRepetetiveGoroutineCreator(&wg)

	updatingGoroutine := repetetiveGoroutineCreator(configs.PollInterval, func() {
		mutex.Lock()
		defer mutex.Unlock()
		metricsCollectionService.UpdateMetrics()
	})

	sendingGoroutine := repetetiveGoroutineCreator(configs.ReportInterval, func() {
		mutex.RLock()
		defer mutex.RUnlock()
		for _, metric := range metricsCollectionService.CollectedMetrics() {
			metricDTO := api.MapDomainMetricToMetricDTO(metric)
			metricJSONData, marshalingError := json.Marshal(metricDTO)
			if marshalingError == nil {
				response, error := http.Post(updateURL, "application/json", bytes.NewBuffer(metricJSONData))
				defer func() {
					if response != nil && response.Body != nil && error == nil {
						response.Body.Close()
					}
				}()
			}
		}
	})

	go updatingGoroutine()
	go sendingGoroutine()

	wg.Wait()
}

func newRepetetiveGoroutineCreator(wg *sync.WaitGroup) func(intervalInSeconds uint, callback Callback) Callback {
	wg.Add(1)
	return func(intervalInSeconds uint, callback Callback) Callback {
		return func() {
			defer wg.Done()
			for {
				time.Sleep(time.Duration(intervalInSeconds) * time.Second)
				callback()
			}
		}
	}
}
