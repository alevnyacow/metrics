package datalayer

type MetricType string

const (
	GAUGE_METRIC   MetricType = "GAUGE"
	COUNTER_METRIC MetricType = "COUNTER"
)

type GaugeMetricName string
type GaugeMetricValue float64

type CounterMetricName string
type CounterMetricValue int64

// Inner structure of memory storage struct.
type MemStorage struct {
	// Map of gauge metrics with keys of gauge metric names
	// and values of gauge metric values.
	gauges map[GaugeMetricName]GaugeMetricValue
	// Map of counter metrics with keys of counter metric names
	// and values of counter metric values.
	counters map[CounterMetricName]CounterMetricValue
}

// Interface of metrics data layer.
type MetricsDataLayer interface {
	// If Counters contained record with given key, its value
	// will be summed with given value. Otherwise, new
	// record will be generated in Counters with given key
	// and given value.
	AddCounterMetric(key CounterMetricName, value CounterMetricValue) (success bool)
	// Returns counter metric value and status flag showing if
	// value was found.
	GetCounterMetricValue(key CounterMetricName) (value CounterMetricValue, wasFound bool)
	// Returns gauge metric value and status flag showing if
	// value was found.
	GetGaugeMetricValue(key GaugeMetricName) (value GaugeMetricValue, wasFound bool)
	// If Gauges contained record with given key, its value
	// will be rewritten with given value. Otherwise, new
	// record will be generated in Gauges with given key
	// and given value.
	SetGaugeMetric(key GaugeMetricName, value GaugeMetricValue) (success bool)
}

var _ MetricsDataLayer = (*MemStorage)(nil)
