package generator

import (
	"github.com/alevnyacow/metrics/internal/datalayer"
	"github.com/alevnyacow/metrics/internal/utils"
)

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

var _ utils.WithLinks = (*Counters)(nil)
var _ utils.WithLinks = (*Gauges)(nil)
