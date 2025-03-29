package memstorage

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
)

func (memstorage *MemStorage) GetCounterValue(key datalayer.CounterName) (value datalayer.CounterValue, wasFound bool) {
	value, wasFound = memstorage.counters[key]
	return
}

func (memstorage *MemStorage) AddCounterMetric(key datalayer.CounterName, value datalayer.CounterValue) (success bool) {
	success = value > 0
	if !success {
		return
	}

	oldValue, foundMetricValue := memstorage.counters[key]

	if !foundMetricValue {
		memstorage.counters[key] = value
		return
	}

	memstorage.counters[key] = oldValue + value
	return
}

func (memstorage *MemStorage) AllCounters() (dtos []datalayer.CounterDTO) {
	dtos = make([]datalayer.CounterDTO, 0)
	for name, value := range memstorage.counters {
		dtos = append(
			dtos,
			datalayer.CounterDTO{
				Name:  name,
				Value: value,
			})
	}
	return
}
