package main

import (
	"github.com/alevnyacow/metrics/internal/generator"
	"github.com/alevnyacow/metrics/internal/utils"
)

func newGeneratorCallback(counters *generator.Counters, gauges *generator.Gauges) func() {
	return func() {
		*counters = generator.GenerateCounters()
		*gauges = generator.GenerateGauges()
	}
}

func newSenderCallback(apiRoot string, counters utils.WithLinks, gauges utils.WithLinks) func() {
	return func() {
		countersLinks := counters.Links(apiRoot)
		gaugesLinks := gauges.Links(apiRoot)

		for _, counterLink := range countersLinks {
			utils.SendPost(counterLink)
		}

		for _, gaugeLink := range gaugesLinks {
			utils.SendPost(gaugeLink)
		}
	}
}

func main() {
	counterMetrics := &generator.Counters{}
	gaugeMetrics := &generator.Gauges{}

	apiRoot := "http://localhost:8080"
	pollInterval := 2
	reportInterval := 10

	generatorCallback := newGeneratorCallback(counterMetrics, gaugeMetrics)
	senderCallback := newSenderCallback(apiRoot, counterMetrics, gaugeMetrics)

	go utils.InfiniteRepetitiveCall(pollInterval, generatorCallback)()
	go utils.InfiniteRepetitiveCall(reportInterval, senderCallback)()

	select {}
}
