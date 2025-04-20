package domain

type MetricType string

const (
	GaugeMetricType   MetricType = "GAUGE"
	CounterMetricType MetricType = "COUNTER"
)

// Metric is a common metric model. It can be
// obtained from Gauge and Counter models. Metric
// can be serialized in JSON.
type Metric struct {
	Name  string     `json:"name"`
	Value string     `json:"value"`
	Type  MetricType `json:"type"`
}

func (metric Metric) IsGauge() bool {
	return metric.Type == GaugeMetricType
}

func (metric Metric) IsCounter() bool {
	return metric.Type == CounterMetricType
}
