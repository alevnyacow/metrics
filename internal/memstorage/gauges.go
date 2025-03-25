package memstorage

import "github.com/alevnyacow/metrics/internal/datalayer"

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
