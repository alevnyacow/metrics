package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/alevnyacow/metrics/internal/api"
	"github.com/alevnyacow/metrics/internal/config"
	"github.com/alevnyacow/metrics/internal/retries"
	"github.com/alevnyacow/metrics/internal/services"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type Callback func()

var client *resty.Client

func init() {
	client = resty.New()
	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() == http.StatusTooManyRequests
		},
	)
}

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
		metricDTOs := make([]api.Metric, 0)
		for _, metric := range metricsCollectionService.CollectedMetrics() {
			metricDTOs = append(metricDTOs, api.MapDomainMetricToMetricDTO(metric))
		}
		metricJSONData, marshalingError := json.Marshal(metricDTOs)
		if marshalingError != nil {
			log.Err(marshalingError).Msg("Error while marshaling metrics DTO")
			return
		}
		requestErr := retries.WithRetries(func() error { return sendPostWithGZippedBody(updateURL, metricJSONData) })
		if requestErr != nil {
			log.Err(requestErr).Msg("Could not send metric update request")
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

func gzippedBytes(data []byte) ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func sendPostWithGZippedBody(url string, body []byte) error {
	gzippedData, gzipError := gzippedBytes(body)
	if gzipError != nil {
		return gzipError
	}
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Content-Encoding", "gzip").
		SetHeader("Accept-Encoding", "gzip").
		SetBody(gzippedData).
		Post(url)
	return err
}
