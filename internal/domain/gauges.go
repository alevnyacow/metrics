package domain

import "strconv"

type GaugeName string
type GaugeValue float64
type GaugeRawValue string
type GaugeRawFloatValue float64

// Gauge model represents gauge metric. Can be
// converted to common Metric model.
type Gauge struct {
	Name  GaugeName
	Value GaugeValue
}

// ToValue converts raw string gauge value to actual
// gauge value. Value must be positive.
func (rawValue GaugeRawValue) ToValue() (value GaugeValue, parsed bool) {
	floatValue, parsingError := strconv.ParseFloat(string(rawValue), 64)
	if parsingError != nil || value < 0 {
		parsed = false
		return
	}
	parsed = true
	value = GaugeValue(floatValue)
	return
}

func (rawFloatValue GaugeRawFloatValue) ToValue() (value GaugeValue, parsed bool) {
	if rawFloatValue < 0 {
		parsed = false
		return
	}
	parsed = true
	value = GaugeValue(rawFloatValue)
	return
}

func (value GaugeValue) ToString() string {
	return strconv.FormatFloat(float64(value), 'f', -1, 64)
}

func (dto Gauge) ToMetricModel() Metric {
	return Metric{
		Name:  string(dto.Name),
		Value: dto.Value.ToString(),
		Type:  GaugeMetricType,
	}
}
