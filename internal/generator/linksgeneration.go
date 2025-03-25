package generator

import "github.com/alevnyacow/metrics/internal/datalayer"

func (counterMetrics *Counters) Links(apiRoot string) (links []string) {
	// Mapping of counter metrics data to their request names.
	counterRequestNamesMapping := map[string]datalayer.CounterValue{
		"PollCount": counterMetrics.PollCount,
	}

	counterLinkBuilder := getCounterLinkBuilder(apiRoot)
	links = make([]string, 0)

	for requestName, metricValue := range counterRequestNamesMapping {
		links = append(links, counterLinkBuilder(requestName, metricValue))
	}

	return
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

	gaugeLinkBuilder := getGaugeLinkBuilder(apiRoot)
	links = make([]string, 0)

	for requestName, metricValue := range gaugeRequestNamesMapping {
		links = append(links, gaugeLinkBuilder(requestName, metricValue))
	}

	return
}
