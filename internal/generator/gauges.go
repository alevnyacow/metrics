package generator

import (
	"fmt"
	"math/rand/v2"
	"runtime"

	"github.com/alevnyacow/metrics/internal/api"
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

func (gaugeMetrics *Gauges) Links(apiRoot string) (links []string) {
	// Mapping of gauge metrics data to their request names.
	gaugeRequestNamesMapping := map[string]datalayer.GaugeValue{
		"Alloc":         gaugeMetrics.Alloc,
		"BuckHashSys":   gaugeMetrics.BuckHashSys,
		"Frees":         gaugeMetrics.Frees,
		"GCCPUFraction": gaugeMetrics.GCCPUFraction,
		"GCSys":         gaugeMetrics.GCSys,
		"HeapAlloc":     gaugeMetrics.HeapAlloc,
		"HeapIdle":      gaugeMetrics.HeapIdle,
		"HeapInuse":     gaugeMetrics.HeapInuse,
		"HeapObjects":   gaugeMetrics.HeapObjects,
		"HeapReleased":  gaugeMetrics.HeapReleased,
		"HeapSys":       gaugeMetrics.HeapSys,
		"LastGC":        gaugeMetrics.LastGC,
		"Lookups":       gaugeMetrics.Lookups,
		"MCacheInuse":   gaugeMetrics.MCacheInuse,
		"MCacheSys":     gaugeMetrics.MCacheSys,
		"MSpanInuse":    gaugeMetrics.MSpanInuse,
		"MSpanSys":      gaugeMetrics.MSpanSys,
		"Mallocs":       gaugeMetrics.Mallocs,
		"NextGC":        gaugeMetrics.NextGC,
		"NumForcedGC":   gaugeMetrics.NumForcedGC,
		"NumGC":         gaugeMetrics.NumGC,
		"OtherSys":      gaugeMetrics.OtherSys,
		"PauseTotalNs":  gaugeMetrics.PauseTotalNs,
		"StackInuse":    gaugeMetrics.StackInuse,
		"StackSys":      gaugeMetrics.StackSys,
		"Sys":           gaugeMetrics.Sys,
		"TotalAlloc":    gaugeMetrics.TotalAlloc,
		"RandomValue":   gaugeMetrics.RandomValue,
	}

	links = make([]string, 0)

	for requestName, metricValue := range gaugeRequestNamesMapping {
		link := fmt.Sprintf(
			"%s/%s/%s/%s/%s",
			apiRoot,
			api.UpdateLinkRoot,
			api.GaugeLinkPath,
			requestName,
			datalayer.GaugeValueToString(metricValue),
		)
		links = append(links, link)
	}

	return
}

var _ utils.WithLinks = (*Gauges)(nil)
