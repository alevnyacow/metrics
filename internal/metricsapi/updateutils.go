package metricsapi

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

func getMetricTypeFromRequest(request *http.Request) (metricType datalayer.MetricType, success bool) {
	pathParamToMetricType := map[string]datalayer.MetricType{
		"gauge":   datalayer.GAUGE_METRIC,
		"counter": datalayer.COUNTER_METRIC,
	}

	metricTypeFromPath := request.PathValue("type")
	metricType, success = pathParamToMetricType[metricTypeFromPath]

	return
}

func getMetricNameAndStringValue(request *http.Request) (metricName string, metricValueAsString string) {
	metricName = request.PathValue("name")
	metricValueAsString = request.PathValue("value")
	return
}

func getCounterPayloadFromRequest(request *http.Request) (payload counterPathParams, success bool) {
	counterMetricName, counterValueAsString := getMetricNameAndStringValue(request)
	counterMetricValue, counterValueParsingError := strconv.ParseInt(counterValueAsString, 10, 64)

	isInErrorState := counterMetricName == "" || counterValueAsString == "" || counterValueParsingError != nil
	success = !isInErrorState

	if !success {
		payload = counterPathParams{}
		return
	}

	payload = counterPathParams{
		name:         datalayer.CounterMetricName(counterMetricName),
		valueCounter: datalayer.CounterMetricValue(counterMetricValue),
	}
	return
}

func getGaugePayloadFromRequest(request *http.Request) (payload gaugePathParams, success bool) {
	gaugeMetricName, gaugeValueAsString := getMetricNameAndStringValue(request)
	gaugeMetricValue, gaugeValueParsingError := strconv.ParseFloat(gaugeValueAsString, 64)

	isInErrorState := gaugeMetricName == "" || gaugeValueAsString == "" || gaugeValueParsingError != nil
	success = !isInErrorState

	if !success {
		payload = gaugePathParams{}
		return
	}

	payload = gaugePathParams{
		name:         datalayer.GaugeMetricName(gaugeMetricName),
		valueCounter: datalayer.GaugeMetricValue(gaugeMetricValue),
	}

	return
}
