package generator

import (
	"strconv"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func getGaugeLinkBuilder(apiRoot string) func(name string, value datalayer.GaugeMetricValue) string {
	return func(name string, value datalayer.GaugeMetricValue) string {
		return apiRoot + "/update/gauge/" + name + "/" + strconv.FormatFloat(float64(value), 'f', -1, 64)
	}
}

func getCounterLinkBuilder(apiRoot string) func(name string, value datalayer.CounterMetricValue) string {
	return func(name string, value datalayer.CounterMetricValue) string {
		return apiRoot + "/update/counter/" + name + "/" + strconv.FormatInt(int64(value), 10)
	}
}
