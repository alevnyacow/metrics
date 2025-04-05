package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/services"
)

type Callback func()

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

func sendPost(url string) {
	request, requestErr := http.NewRequest("POST", url, nil)
	if requestErr == nil {
		request.Header.Set("Content-Type", "text/plain")
		client := http.Client{}
		client.Do(request)
	}
}

func main() {
	mutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	configs := config.ParseAgentConfigs()
	metricsCollectorService := services.NewMetricsCollectorService()
	repetetiveGoroutineCreator := newRepetetiveGoroutineCreator(&wg)

	updatingGoroutine := repetetiveGoroutineCreator(configs.PollInterval, func() {
		mutex.Lock()
		defer mutex.Unlock()
		metricsCollectorService.UpdateMetrics()
	})

	sendingGoroutine := repetetiveGoroutineCreator(configs.ReportInterval, func() {
		mutex.RLock()
		defer mutex.RUnlock()
		getUpdateCounterLink, getUpdateGaugeLink := api.MetricUpdateRoutes(configs.APIHost)
		for _, counter := range metricsCollectorService.Counters {
			sendPost(getUpdateCounterLink(counter))
		}
		for _, gauge := range metricsCollectorService.Gauges {
			sendPost(getUpdateGaugeLink(gauge))
		}
	})

	go updatingGoroutine()
	go sendingGoroutine()

	wg.Wait()
}
