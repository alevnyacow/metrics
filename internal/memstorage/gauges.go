package memstorage

func (memstorage *MemStorage) GetGaugeMetricValue(key GaugeMetricName) (value GaugeMetricValue, wasFound bool) {
	value, wasFound = memstorage.gauges[key]
	return
}

// If Gauges contained record with given key, its value
// will be rewritten with given value. Otherwise, new
// record will be generated in Gauges with given key
// and given value.
func (memstorage *MemStorage) SetGaugeMetric(key GaugeMetricName, value GaugeMetricValue) (createdNewGaugeMetric bool) {
	_, foundMetricValue := memstorage.gauges[key]
	createdNewGaugeMetric = !foundMetricValue
	memstorage.gauges[key] = value
	return
}
