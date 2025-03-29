package memstorage

import "github.com/alevnyacow/metrics/internal/datalayer"

// Inner structure of memory storage struct.
type MemStorage struct {
	// Map of gauge metrics with keys of gauge metric names
	// and values of gauge metric values.
	gauges map[datalayer.GaugeName]datalayer.GaugeValue
	// Map of counter metrics with keys of counter metric names
	// and values of counter metric values.
	counters map[datalayer.CounterName]datalayer.CounterValue
}

// Static check if MemStorage implements
// DataLayer interface.
var _ datalayer.DataLayer = (*MemStorage)(nil)
