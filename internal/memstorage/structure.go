package memstorage

type GaugeMetricName string
type GaugeMetricValue float64

type CounterMetricName string
type CounterMetricValue int64

type MemStorage struct {
	// Map of gauge metrics with keys of gauge metric names
	// and values of gauge metric values.
	gauges map[GaugeMetricName]GaugeMetricValue
	// Map of counter metrics with keys of counter metric names
	// and values of counter metric values.
	counters map[CounterMetricName]CounterMetricValue
}
