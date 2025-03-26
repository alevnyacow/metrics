package datalayer

import "strconv"

// Default implementation of getting gauge
// value string representation.
func GaugeValueToString(value GaugeValue) string {
	return strconv.FormatFloat(float64(value), 'f', -1, 64)
}

// Default implementation of getting gauge
// value from raw string value.
func GaugeValueFromString(gaugeValueAsString string) (gaugeValue GaugeValue, valueWasParsed bool) {
	value, parsingError := strconv.ParseFloat(gaugeValueAsString, 64)
	if parsingError != nil {
		valueWasParsed = false
		return
	}
	valueWasParsed = true
	gaugeValue = GaugeValue(value)
	return
}
