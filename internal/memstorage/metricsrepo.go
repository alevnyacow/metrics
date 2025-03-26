package memstorage

import "github.com/alevnyacow/metrics/internal/datalayer"

func (memstorage *MemStorage) AllMetrics() (dtos []datalayer.MetricDTO) {
	dtos = make([]datalayer.MetricDTO, 0)
	for _, counterDTO := range memstorage.AllCounters() {
		dtos = append(dtos, datalayer.CounterDTOToMetricDTO(counterDTO))
	}
	for _, gaugeDTO := range memstorage.AllGauges() {
		dtos = append(dtos, datalayer.GaugeDTOToMetricDTO(gaugeDTO))
	}
	return
}
