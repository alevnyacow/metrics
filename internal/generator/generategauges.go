package generator

import (
	"math/rand/v2"
	"runtime"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func GenerateGauges() Gauges {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return Gauges{
		Alloc:         datalayer.GaugeValue(stats.Alloc),
		BuckHashSys:   datalayer.GaugeValue(stats.BuckHashSys),
		Frees:         datalayer.GaugeValue(stats.Frees),
		GCCPUFraction: datalayer.GaugeValue(stats.GCCPUFraction),
		GCSys:         datalayer.GaugeValue(stats.GCSys),
		HeapAlloc:     datalayer.GaugeValue(stats.HeapAlloc),
		HeapIdle:      datalayer.GaugeValue(stats.HeapIdle),
		HeapInuse:     datalayer.GaugeValue(stats.HeapInuse),
		HeapObjects:   datalayer.GaugeValue(stats.HeapObjects),
		HeapReleased:  datalayer.GaugeValue(stats.HeapReleased),
		HeapSys:       datalayer.GaugeValue(stats.HeapSys),
		LastGC:        datalayer.GaugeValue(stats.LastGC),
		Lookups:       datalayer.GaugeValue(stats.Lookups),
		MCacheInuse:   datalayer.GaugeValue(stats.MCacheInuse),
		MCacheSys:     datalayer.GaugeValue(stats.MCacheSys),
		MSpanInuse:    datalayer.GaugeValue(stats.MSpanInuse),
		MSpanSys:      datalayer.GaugeValue(stats.MSpanSys),
		Mallocs:       datalayer.GaugeValue(stats.Mallocs),
		NextGC:        datalayer.GaugeValue(stats.NextGC),
		NumForcedGC:   datalayer.GaugeValue(stats.NextGC),
		NumGC:         datalayer.GaugeValue(stats.NumGC),
		OtherSys:      datalayer.GaugeValue(stats.OtherSys),
		PauseTotalNs:  datalayer.GaugeValue(stats.PauseTotalNs),
		StackInuse:    datalayer.GaugeValue(stats.StackInuse),
		StackSys:      datalayer.GaugeValue(stats.StackSys),
		Sys:           datalayer.GaugeValue(stats.Sys),
		TotalAlloc:    datalayer.GaugeValue(stats.TotalAlloc),
		RandomValue:   datalayer.GaugeValue(rand.Float64() * 100),
	}
}
