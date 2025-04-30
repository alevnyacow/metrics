package services

import (
	"math/rand/v2"
	"runtime"

	"github.com/alevnyacow/metrics/internal/domain"
	"github.com/rs/zerolog/log"
)

type AgentCounters struct {
	PollCount domain.CounterValue
}

type AgentGauges struct {
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
	RandomValue domain.GaugeValue
}

// MetricsCollectionService provides logic of
// updating and storing collected metrics.
type MetricsCollectionService struct {
	counters []domain.Counter
	gauges   []domain.Gauge
}

func (service *MetricsCollectionService) generateCounters() AgentCounters {
	return AgentCounters{
		PollCount: 1,
	}
}

// generateGauges generate gauges based on runtime.MemStats and
// some random values
func (service *MetricsCollectionService) generateGauges() AgentGauges {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return AgentGauges{
		Alloc:         domain.GaugeValue(stats.Alloc),
		BuckHashSys:   domain.GaugeValue(stats.BuckHashSys),
		Frees:         domain.GaugeValue(stats.Frees),
		GCCPUFraction: domain.GaugeValue(stats.GCCPUFraction),
		GCSys:         domain.GaugeValue(stats.GCSys),
		HeapAlloc:     domain.GaugeValue(stats.HeapAlloc),
		HeapIdle:      domain.GaugeValue(stats.HeapIdle),
		HeapInuse:     domain.GaugeValue(stats.HeapInuse),
		HeapObjects:   domain.GaugeValue(stats.HeapObjects),
		HeapReleased:  domain.GaugeValue(stats.HeapReleased),
		HeapSys:       domain.GaugeValue(stats.HeapSys),
		LastGC:        domain.GaugeValue(stats.LastGC),
		Lookups:       domain.GaugeValue(stats.Lookups),
		MCacheInuse:   domain.GaugeValue(stats.MCacheInuse),
		MCacheSys:     domain.GaugeValue(stats.MCacheSys),
		MSpanInuse:    domain.GaugeValue(stats.MSpanInuse),
		MSpanSys:      domain.GaugeValue(stats.MSpanSys),
		Mallocs:       domain.GaugeValue(stats.Mallocs),
		NextGC:        domain.GaugeValue(stats.NextGC),
		NumForcedGC:   domain.GaugeValue(stats.NextGC),
		NumGC:         domain.GaugeValue(stats.NumGC),
		OtherSys:      domain.GaugeValue(stats.OtherSys),
		PauseTotalNs:  domain.GaugeValue(stats.PauseTotalNs),
		StackInuse:    domain.GaugeValue(stats.StackInuse),
		StackSys:      domain.GaugeValue(stats.StackSys),
		Sys:           domain.GaugeValue(stats.Sys),
		TotalAlloc:    domain.GaugeValue(stats.TotalAlloc),
		RandomValue:   domain.GaugeValue(rand.Float64() * 100),
	}
}

// UpdateMetrics generates actual metrics and
// writes obtained values in service.
func (service *MetricsCollectionService) UpdateMetrics() {
	counters := service.generateCounters()
	gauges := service.generateGauges()
	service.counters = counters.toMetrics()
	log.Info().Msg("Collected and updated counters")
	service.gauges = gauges.toMetrics()
	log.Info().Msg("Collected and updated gauges")
}

func (service *MetricsCollectionService) CollectedMetrics() []domain.Metric {
	metrics := make([]domain.Metric, 0)
	for _, counter := range service.counters {
		metrics = append(metrics, counter.ToMetricModel())
	}
	for _, gauge := range service.gauges {
		metrics = append(metrics, gauge.ToMetricModel())
	}
	return metrics
}

func (counters AgentCounters) toMetrics() []domain.Counter {
	result := make([]domain.Counter, 0)
	result = append(result, domain.Counter{Name: "PollCount", Value: counters.PollCount})
	return result
}

func (gauges AgentGauges) toMetrics() []domain.Gauge {
	result := make([]domain.Gauge, 0)
	result = append(result, domain.Gauge{Name: "Alloc", Value: gauges.Alloc})
	result = append(result, domain.Gauge{Name: "BuckHashSys", Value: gauges.BuckHashSys})
	result = append(result, domain.Gauge{Name: "Frees", Value: gauges.Frees})
	result = append(result, domain.Gauge{Name: "GCCPUFraction", Value: gauges.GCCPUFraction})
	result = append(result, domain.Gauge{Name: "GCSys", Value: gauges.GCSys})
	result = append(result, domain.Gauge{Name: "HeapAlloc", Value: gauges.HeapAlloc})
	result = append(result, domain.Gauge{Name: "HeapIdle", Value: gauges.HeapIdle})
	result = append(result, domain.Gauge{Name: "HeapInuse", Value: gauges.HeapInuse})
	result = append(result, domain.Gauge{Name: "HeapObjects", Value: gauges.HeapObjects})
	result = append(result, domain.Gauge{Name: "HeapReleased", Value: gauges.HeapReleased})
	result = append(result, domain.Gauge{Name: "HeapSys", Value: gauges.HeapSys})
	result = append(result, domain.Gauge{Name: "LastGC", Value: gauges.LastGC})
	result = append(result, domain.Gauge{Name: "Lookups", Value: gauges.Lookups})
	result = append(result, domain.Gauge{Name: "MCacheInuse", Value: gauges.MCacheInuse})
	result = append(result, domain.Gauge{Name: "MCacheSys", Value: gauges.MCacheSys})
	result = append(result, domain.Gauge{Name: "MSpanInuse", Value: gauges.MSpanInuse})
	result = append(result, domain.Gauge{Name: "MSpanSys", Value: gauges.MSpanSys})
	result = append(result, domain.Gauge{Name: "Mallocs", Value: gauges.Mallocs})
	result = append(result, domain.Gauge{Name: "NextGC", Value: gauges.NextGC})
	result = append(result, domain.Gauge{Name: "NumForcedGC", Value: gauges.NumForcedGC})
	result = append(result, domain.Gauge{Name: "NumGC", Value: gauges.NumGC})
	result = append(result, domain.Gauge{Name: "OtherSys", Value: gauges.OtherSys})
	result = append(result, domain.Gauge{Name: "PauseTotalNs", Value: gauges.PauseTotalNs})
	result = append(result, domain.Gauge{Name: "StackInuse", Value: gauges.StackInuse})
	result = append(result, domain.Gauge{Name: "StackSys", Value: gauges.StackSys})
	result = append(result, domain.Gauge{Name: "Sys", Value: gauges.Sys})
	result = append(result, domain.Gauge{Name: "TotalAlloc", Value: gauges.TotalAlloc})
	result = append(result, domain.Gauge{Name: "RandomValue", Value: gauges.RandomValue})
	return result
}
