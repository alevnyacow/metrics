package metricsgenerator

import "github.com/alevnyacow/metrics/internal/datalayer"

type GaugeMetrics struct {
	Alloc,
	BuckHashSys,
	Frees,
	GCCPUFraction,
	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	Mallocs,
	NextGC,
	NumForcedGC,
	NumGC,
	OtherSys,
	PauseTotalNs,
	StackInuse,
	StackSys,
	Sys,
	TotalAlloc,
	RandomValue datalayer.GaugeMetricValue
}

type CounterMetrics struct {
	PollCount datalayer.CounterMetricValue
}

type WithLinksGeneration interface {
	Links(apiRoot string) (links []string)
}

var _ WithLinksGeneration = (*CounterMetrics)(nil)
var _ WithLinksGeneration = (*GaugeMetrics)(nil)
