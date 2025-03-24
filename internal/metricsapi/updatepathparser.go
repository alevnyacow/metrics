package metricsapi

// Path parsing utility helpers for
// update POST request.

import (
	"net/http"
	"strconv"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

type counterPathParams struct {
	name         datalayer.CounterMetricName
	valueCounter datalayer.CounterMetricValue
}

type gaugePathParams struct {
	name         datalayer.GaugeMetricName
	valueCounter datalayer.GaugeMetricValue
}

type updatePathParsingResult struct {
	parsedName  bool
	parsedValue bool
}

func parseMetricTypeFromRequest(request *http.Request) (metricType datalayer.MetricType, success bool) {
	pathParamToMetricType := map[string]datalayer.MetricType{
		"gauge":   datalayer.GaugeMetricType,
		"counter": datalayer.CounterMetricType,
	}

	metricTypeFromPath := request.PathValue("type")
	metricType, success = pathParamToMetricType[metricTypeFromPath]

	return
}

func parseMetricNameAndStringValue(request *http.Request) (metricName string, metricValueAsString string) {
	metricName = request.PathValue("name")
	metricValueAsString = request.PathValue("value")
	return
}

func parseCounterPayloadFromRequest(request *http.Request) (payload counterPathParams, result updatePathParsingResult) {
	counterMetricName, counterValueAsString := parseMetricNameAndStringValue(request)
	counterMetricValue, counterValueParsingError := strconv.ParseInt(counterValueAsString, 10, 64)

	result = updatePathParsingResult{
		parsedName:  counterMetricName != "",
		parsedValue: counterValueParsingError == nil,
	}

	payload = counterPathParams{
		name:         datalayer.CounterMetricName(counterMetricName),
		valueCounter: datalayer.CounterMetricValue(counterMetricValue),
	}

	return
}

func parseGaugePayloadFromRequest(request *http.Request) (payload gaugePathParams, result updatePathParsingResult) {
	gaugeMetricName, gaugeValueAsString := parseMetricNameAndStringValue(request)
	gaugeMetricValue, gaugeValueParsingError := strconv.ParseFloat(gaugeValueAsString, 64)

	result = updatePathParsingResult{
		parsedName:  gaugeMetricName != "",
		parsedValue: gaugeValueParsingError == nil,
	}

	payload = gaugePathParams{
		name:         datalayer.GaugeMetricName(gaugeMetricName),
		valueCounter: datalayer.GaugeMetricValue(gaugeMetricValue),
	}

	return
}
