package memstorage

import "github.com/alevnyacow/metrics/internal/datalayer"

// Returns prepared MemStorage instanse.
func NewMemStorage() (instanse *MemStorage) {
	instanse = &MemStorage{
		gauges:   make(map[datalayer.GaugeName]datalayer.GaugeValue),
		counters: make(map[datalayer.CounterName]datalayer.CounterValue),
	}

	return
}
