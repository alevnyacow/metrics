package datalayer

func GaugeDTOToMetricDTO(gaugeDTO GaugeDTO) MetricDTO {
	return MetricDTO(gaugeDTO)
}

func CounterDTOToMetricDTO(counterDTO CounterDTO) MetricDTO {
	return MetricDTO(counterDTO)
}
