package generator

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
)

func getGaugeLinkBuilder(apiRoot string) func(name string, value datalayer.GaugeValue) string {
	return func(name string, value datalayer.GaugeValue) string {
		return apiRoot + "/update/gauge/" + name + "/" + datalayer.GaugeValueToString(value)
	}
}

func getCounterLinkBuilder(apiRoot string) func(name string, value datalayer.CounterValue) string {
	return func(name string, value datalayer.CounterValue) string {
		return apiRoot + "/update/counter/" + name + "/" + datalayer.CounterValueToString(value)
	}
}
