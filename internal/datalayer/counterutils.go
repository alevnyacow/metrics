package datalayer

import "strconv"

func CounterValueToString(value CounterValue) string {
	return strconv.FormatInt(int64(value), 10)
}

// Parse counter value data from raw string data.
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
