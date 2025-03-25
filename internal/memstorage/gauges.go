package memstorage

import (
	"strconv"

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

func (memstorage *MemStorage) AllGauges() (dtos []datalayer.MetricDTO) {
	dtos = make([]datalayer.MetricDTO, 0)
	for name, value := range memstorage.gauges {
		dtos = append(
			dtos,
			datalayer.MetricDTO{
				Name:  string(name),
				Value: strconv.FormatFloat(float64(value), 'f', -1, 64),
			})
	}
	return
}
