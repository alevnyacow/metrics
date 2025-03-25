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
	RandomValue datalayer.GaugeValue
}

type Counters struct {
	PollCount datalayer.CounterValue
}

type WithLinks interface {
	Links(apiRoot string) (links []string)
}

var _ WithLinks = (*Counters)(nil)
var _ WithLinks = (*Gauges)(nil)
