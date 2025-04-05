package domain

type MetricType string

const (
	GaugeMetricType   MetricType = "GAUGE"
	CounterMetricType MetricType = "COUNTER"
)

// Metric is a common metric model. It can be
// obtained from Gauge and Metric models. Metric
// can be serialized in JSON.
type Metric struct {
	Name  string     `json:"name"`
	Value string     `json:"value"`
	Type  MetricType `json:"type"`
}
