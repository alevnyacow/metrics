package datalayer

type MetricType string

const (
	GaugeMetricType   MetricType = "GAUGE"
	CounterMetricType MetricType = "COUNTER"
)

type GaugeName string
type GaugeValue float64

type CounterName string
type CounterValue int64

// Inner structure of memory storage struct.
type MemStorage struct {
	// Map of gauge metrics with keys of gauge metric names
	// and values of gauge metric values.
	gauges map[GaugeName]GaugeValue
	// Map of counter metrics with keys of counter metric names
	// and values of counter metric values.
	counters map[CounterName]CounterValue
}

// Interface of metrics data layer.
type MetricsDataLayer interface {
	// If Counters contained record with given key, its value
	// will be summed with given value. Otherwise, new
	// record will be generated in Counters with given key
	// and given value.
	AddCounterMetric(key CounterName, value CounterValue) (success bool)
	// Returns counter metric value and status flag showing if
	// value was found.
	GetCounterValue(key CounterName) (value CounterValue, wasFound bool)
	// Returns gauge metric value and status flag showing if
	// value was found.
	GetGaugeValue(key GaugeName) (value GaugeValue, wasFound bool)
	// If Gauges contained record with given key, its value
	// will be rewritten with given value. Otherwise, new
	// record will be generated in Gauges with given key
	// and given value.
	SetGaugeMetric(key GaugeName, value GaugeValue) (success bool)
}
