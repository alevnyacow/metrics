package datalayer

// Definite metric type.
type MetricType string

const (
	// Gauge metric.
	GaugeMetricType MetricType = "GAUGE"
	// Counter metric.
	CounterMetricType MetricType = "COUNTER"
)

// Metric Data Transfer Object struct, which
// is general metric representation. Can be
// serialized into JSON.
type MetricDTO struct {
	Name  string     `json:"name"`
	Value string     `json:"value"`
	Type  MetricType `json:"type"`
}

// Counter metric name type.
type CounterName string

// Counter metric value type.
type CounterValue int64

// Counter metric Data Transfer Object.
type CounterDTO struct {
	Name  CounterName
	Value CounterValue
}

// Gauge metric name type.
type GaugeName string

// Gauge metric value type.
type GaugeValue float64

// Gauge metric Data Transfere Object.
type GaugeDTO struct {
	Name  GaugeName
	Value GaugeValue
}

// Interface of counters repository.
type CountersRepository interface {
	// If Counters contained record with given key, its value
	// will be summed with given value. Otherwise, new
	// record will be generated in Counters with given key
	// and given value.
	AddCounterMetric(key CounterName, value CounterValue) (success bool)
	// Returns counter metric value and status flag showing if
	// value was found.
	GetCounterValue(key CounterName) (value CounterValue, wasFound bool)
	// Returns information about all counters.
	AllCounters() []CounterDTO
}

// Interface of gauges repository.
type GaugesRepository interface {
	// Returns gauge metric value and status flag showing if
	// value was found.
	GetGaugeValue(key GaugeName) (value GaugeValue, wasFound bool)
	// If Gauges contained record with given key, its value
	// will be rewritten with given value. Otherwise, new
	// record will be generated in Gauges with given key
	// and given value.
	SetGaugeMetric(key GaugeName, value GaugeValue) (success bool)
	// Returns information about all gauges.
	AllGauges() []GaugeDTO
}

// Interface of common metrics repository
// for all metrics regardless its type.
type MetricsRepository interface {
	AllMetrics() []MetricDTO
}

// Interface of metrics data layer.
type DataLayer interface {
	CountersRepository
	GaugesRepository
	MetricsRepository
}
