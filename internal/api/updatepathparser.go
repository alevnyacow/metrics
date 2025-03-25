package api

// Path parsing utility helpers for
// update POST request.

import (
	"net/http"
	"strconv"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

type counterPathParams struct {
	name         datalayer.CounterName
	valueCounter datalayer.CounterValue
}

type gaugePathParams struct {
	name         datalayer.GaugeName
	valueCounter datalayer.GaugeValue
}

type updatePathParsingResult struct {
	parsedName  bool
	parsedValue bool
}

type metric string

const (
	gaugeMetricType   metric = "GAUGE"
	counterMetricType metric = "COUNTER"
)

func parseMetricTypeFromRequest(request *http.Request) (metricType metric, success bool) {
	pathParamToMetricType := map[string]metric{
		"gauge":   gaugeMetricType,
		"counter": counterMetricType,
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
	CounterName, counterValueAsString := parseMetricNameAndStringValue(request)
	CounterValue, counterValueParsingError := strconv.ParseInt(counterValueAsString, 10, 64)

	result = updatePathParsingResult{
		parsedName:  CounterName != "",
		parsedValue: counterValueParsingError == nil,
	}

	payload = counterPathParams{
		name:         datalayer.CounterName(CounterName),
		valueCounter: datalayer.CounterValue(CounterValue),
	}

	return
}

func parseGaugePayloadFromRequest(request *http.Request) (payload gaugePathParams, result updatePathParsingResult) {
	GaugeName, gaugeValueAsString := parseMetricNameAndStringValue(request)
	GaugeValue, gaugeValueParsingError := strconv.ParseFloat(gaugeValueAsString, 64)

	result = updatePathParsingResult{
		parsedName:  GaugeName != "",
		parsedValue: gaugeValueParsingError == nil,
	}

	payload = gaugePathParams{
		name:         datalayer.GaugeName(GaugeName),
		valueCounter: datalayer.GaugeValue(GaugeValue),
	}

	return
}
