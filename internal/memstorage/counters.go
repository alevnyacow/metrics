package memstorage

func (memstorage *MemStorage) GetCounterMetricValue(key CounterMetricName) (value CounterMetricValue, wasFound bool) {
	value, wasFound = memstorage.counters[key]
	return
}

// If Counters contained record with given key, its value
// will be summed with given value. Otherwise, new
// record will be generated in Counters with given key
// and given value.
func (memstorage *MemStorage) AddCounterMetric(key CounterMetricName, value CounterMetricValue) (createdNewCounterMetric bool) {
	oldValue, foundMetricValue := memstorage.counters[key]
	createdNewCounterMetric = !foundMetricValue

	if createdNewCounterMetric {
		memstorage.counters[key] = value
		return
	}

	memstorage.counters[key] = oldValue + value
	return
}
