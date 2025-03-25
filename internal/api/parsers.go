package api

import (
	"net/http"
	"strconv"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

// Definite metric type.
type metric string

const (
	// Gauge metric.
	gaugeMetricType metric = "GAUGE"
	// Counter metric.
	counterMetricType metric = "COUNTER"
)

// Parse data from "type" path parameter and return
// definite metric type or nothing if we could not map
// this result from path parameter value.
func parseMetricType(request *http.Request) (metricType metric, success bool) {
	pathParamToMetricType := map[string]metric{
		"gauge":   gaugeMetricType,
		"counter": counterMetricType,
	}
	metricTypeFromPath := request.PathValue("type")
	metricType, success = pathParamToMetricType[metricTypeFromPath]
	return
}

// Parse string data from "name" path parameter.
func parseMetricName(request *http.Request) (metricName string, nameWasProvided bool) {
	metricName = request.PathValue("name")
	nameWasProvided = metricName != ""
	return
}

// Parse string data from "value" path parameter.
func parseStringValue(request *http.Request) (valueAsString string, valueWasProvided bool) {
	valueAsString = request.PathValue("value")
	valueWasProvided = valueAsString != ""
	return
}

// Parse counter value data from raw string data.
func parseCounterValue(counterValueAsString string) (counterValue datalayer.CounterValue, valueWasParsed bool) {
	value, parsingError := strconv.ParseInt(counterValueAsString, 10, 64)
	if parsingError != nil {
		valueWasParsed = false
		return
	}
	valueWasParsed = true
	counterValue = datalayer.CounterValue(value)
	return
}

// Parse gauge value data from raw string data.
func parseGaugeValue(gaugeValueAsString string) (gaugeValue datalayer.GaugeValue, valueWasParsed bool) {
	value, parsingError := strconv.ParseFloat(gaugeValueAsString, 64)
	if parsingError != nil {
		valueWasParsed = false
		return
	}
	valueWasParsed = true
	gaugeValue = datalayer.GaugeValue(value)
	return
}
