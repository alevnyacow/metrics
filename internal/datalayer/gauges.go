package datalayer

func (memstorage *MemStorage) GetGaugeMetricValue(key GaugeMetricName) (value GaugeMetricValue, wasFound bool) {
	value, wasFound = memstorage.gauges[key]
	return
}

func (memstorage *MemStorage) SetGaugeMetric(key GaugeMetricName, value GaugeMetricValue) (success bool) {
	success = value > 0
	if success {
		memstorage.gauges[key] = value
	}
	return
}
