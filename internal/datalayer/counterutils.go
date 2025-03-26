package datalayer

import "strconv"

// Default implementation of getting counter
// value string representation.
func CounterValueToString(value CounterValue) string {
	return strconv.FormatInt(int64(value), 10)
}

// Default implementation of getting counter
// value from raw string value.
func CounterValueFromString(counterValueAsString string) (counterValue CounterValue, valueWasParsed bool) {
	value, parsingError := strconv.ParseInt(counterValueAsString, 10, 64)
	if parsingError != nil {
		valueWasParsed = false
		return
	}
	valueWasParsed = true
	counterValue = CounterValue(value)
	return
}
