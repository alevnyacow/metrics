package memstorage

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
)

func (memstorage *MemStorage) GetGaugeValue(key datalayer.GaugeName) (value datalayer.GaugeValue, wasFound bool) {
	value, wasFound = memstorage.gauges[key]
	return
}

func (memstorage *MemStorage) SetGaugeMetric(key datalayer.GaugeName, value datalayer.GaugeValue) (success bool) {
	success = value > 0
	if success {
		memstorage.gauges[key] = value
	}
	return
}

func (memstorage *MemStorage) AllGauges() (dtos []datalayer.GaugeDTO) {
	dtos = make([]datalayer.GaugeDTO, 0)
	for name, value := range memstorage.gauges {
		dtos = append(
			dtos,
			datalayer.GaugeDTO{
				Name:  string(name),
				Value: datalayer.GaugeValueToString(value),
			})
	}
	return
}
