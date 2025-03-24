package datalayer

func (memstorage *MemStorage) GetCounterMetricValue(key CounterMetricName) (value CounterMetricValue, wasFound bool) {
	value, wasFound = memstorage.counters[key]
	return
}

func (memstorage *MemStorage) AddCounterMetric(key CounterMetricName, value CounterMetricValue) (success bool) {
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
