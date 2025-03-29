package datalayer

func MapCounterDTOToMetricDTO(counter CounterDTO) MetricDTO {
	return MetricDTO{
		Name:  string(counter.Name),
		Value: CounterValueToString(counter.Value),
		Type:  CounterMetricType,
	}
}

func MapGaugeDTOToMetricDTO(gauge GaugeDTO) MetricDTO {
	return MetricDTO{
		Name:  string(gauge.Name),
		Value: GaugeValueToString(gauge.Value),
		Type:  GaugeMetricType,
	}
}
