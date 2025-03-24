package main

import (
	"fmt"
	"time"

	"github.com/alevnyacow/metrics/internal/metricsgenerator"
)

func generateMetrics(intervalInSeconds int64, counters *metricsgenerator.CounterMetrics, gauges *metricsgenerator.GaugeMetrics) {
	for {
		time.Sleep(time.Duration(intervalInSeconds) * time.Second)
		*counters = metricsgenerator.GenerateCounters()
		*gauges = metricsgenerator.GenerateGauges()
	}
}

func sendMetrics(intervalInSeconds int64, apiRoot string, counters metricsgenerator.WithLinksGeneration, gauges metricsgenerator.WithLinksGeneration) {
	for {
		time.Sleep(time.Duration(intervalInSeconds) * time.Second)

		countersLinks := counters.Links(apiRoot)
		gaugesLinks := gauges.Links(apiRoot)

		for _, counterLink := range countersLinks {
			fmt.Println("COUNTER LINK", counterLink)
		}

		for _, gaugeLink := range gaugesLinks {
			fmt.Println("GAUGE LINK", gaugeLink)
		}
	}
}

func main() {
	counterMetrics := &metricsgenerator.CounterMetrics{}
	gaugeMetrics := &metricsgenerator.GaugeMetrics{}

	go generateMetrics(2, counterMetrics, gaugeMetrics)
	go sendMetrics(5, "http://localhost:8080", counterMetrics, gaugeMetrics)

	select {}
}
