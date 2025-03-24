package datalayer

// Returns prepared MemStorage instanse.
func NewMemStorage() (instanse *MemStorage) {
	instanse = &MemStorage{
		gauges:   make(map[GaugeMetricName]GaugeMetricValue),
		counters: make(map[CounterMetricName]CounterMetricValue),
	}

	return
}
