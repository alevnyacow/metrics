package domain

import "strconv"

type CounterName string
type CounterValue int64
type CounterRawValue string

// Counter model represents counter metric. Can be
// converted to common Metric model.
type Counter struct {
	Name  CounterName
	Value CounterValue
}

func (value CounterValue) ToString() string {
	return strconv.FormatInt(int64(value), 10)
}

// ToValue converts raw string counter value to actual
// counter value. Value must be positive.
func (rawValue CounterRawValue) ToValue() (value CounterValue, parsed bool) {
	intValue, parsingError := strconv.ParseInt(string(rawValue), 10, 64)
	if parsingError != nil {
		parsed = false
		return
	}
	parsed = true
	value = CounterValue(intValue)
	return
}

func (dto Counter) ToMetricModel() Metric {
	return Metric{
		Name:  string(dto.Name),
		Value: dto.Value.ToString(),
		Type:  CounterMetricType,
	}
}
