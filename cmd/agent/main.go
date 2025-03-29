package main

import (
	"github.com/alevnyacow/metrics/internal/config"
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
	configs := config.ParseAgentConfigs()
	counterMetrics := &generator.Counters{}
	gaugeMetrics := &generator.Gauges{}

	generatorCallback := newGeneratorCallback(counterMetrics, gaugeMetrics)
	senderCallback := newSenderCallback(configs.APIHost, counterMetrics, gaugeMetrics)

	go utils.InfiniteRepetitiveCall(configs.PollInterval, generatorCallback)()
	go utils.InfiniteRepetitiveCall(configs.ReportInterval, senderCallback)()
	select {}
}
