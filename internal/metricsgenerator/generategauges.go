package metricsgenerator

import (
	"math/rand/v2"
	"runtime"

	"github.com/alevnyacow/metrics/internal/datalayer"
)

func GenerateGauges() GaugeMetrics {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return GaugeMetrics{
		Alloc:         datalayer.GaugeMetricValue(stats.Alloc),
		BuckHashSys:   datalayer.GaugeMetricValue(stats.BuckHashSys),
		Frees:         datalayer.GaugeMetricValue(stats.Frees),
		GCCPUFraction: datalayer.GaugeMetricValue(stats.GCCPUFraction),
		GCSys:         datalayer.GaugeMetricValue(stats.GCSys),
		HeapAlloc:     datalayer.GaugeMetricValue(stats.HeapAlloc),
		HeapIdle:      datalayer.GaugeMetricValue(stats.HeapIdle),
		HeapInuse:     datalayer.GaugeMetricValue(stats.HeapInuse),
		HeapObjects:   datalayer.GaugeMetricValue(stats.HeapObjects),
		HeapReleased:  datalayer.GaugeMetricValue(stats.HeapReleased),
		HeapSys:       datalayer.GaugeMetricValue(stats.HeapSys),
		LastGC:        datalayer.GaugeMetricValue(stats.LastGC),
		Lookups:       datalayer.GaugeMetricValue(stats.Lookups),
		MCacheInuse:   datalayer.GaugeMetricValue(stats.MCacheInuse),
		MCacheSys:     datalayer.GaugeMetricValue(stats.MCacheSys),
		MSpanInuse:    datalayer.GaugeMetricValue(stats.MSpanInuse),
		MSpanSys:      datalayer.GaugeMetricValue(stats.MSpanSys),
		Mallocs:       datalayer.GaugeMetricValue(stats.Mallocs),
		NextGC:        datalayer.GaugeMetricValue(stats.NextGC),
		NumForcedGC:   datalayer.GaugeMetricValue(stats.NextGC),
		NumGC:         datalayer.GaugeMetricValue(stats.NumGC),
		OtherSys:      datalayer.GaugeMetricValue(stats.OtherSys),
		PauseTotalNs:  datalayer.GaugeMetricValue(stats.PauseTotalNs),
		StackInuse:    datalayer.GaugeMetricValue(stats.StackInuse),
		StackSys:      datalayer.GaugeMetricValue(stats.StackSys),
		Sys:           datalayer.GaugeMetricValue(stats.Sys),
		TotalAlloc:    datalayer.GaugeMetricValue(stats.TotalAlloc),
		RandomValue:   datalayer.GaugeMetricValue(rand.Float64() * 100),
	}
}
