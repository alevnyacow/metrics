package datalayer

type GaugeName string
type GaugeValue float64

type CounterName string
type CounterValue int64

type MetricDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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
	AllCounters() []MetricDTO
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
	AllGauges() []MetricDTO
}

// Interface of metrics data layer.
type DataLayer interface {
	CountersRepository
	GaugesRepository
}
