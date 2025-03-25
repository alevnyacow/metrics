package generator

import "github.com/alevnyacow/metrics/internal/datalayer"

type Gauges struct {
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

type Counters struct {
	PollCount datalayer.CounterMetricValue
}

type WithLinksGeneration interface {
	Links(apiRoot string) (links []string)
}

var _ WithLinksGeneration = (*Counters)(nil)
var _ WithLinksGeneration = (*Gauges)(nil)
