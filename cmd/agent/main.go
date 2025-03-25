package main

import (
	"fmt"

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
		// gaugesLinks := gauges.Links(apiRoot)

		for _, counterLink := range countersLinks {
			fmt.Println("CALLED " + counterLink)
			// req, g := http.NewRequest("POST", counterLink, nil)
			// if g != nil {
			// 	fmt.Println(g.Error())
			// }
			// client := http.Client{
			// 	Timeout: 10 * time.Second, // Set timeout to 10 seconds
			// }
			// client.Do(req)

			/**
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Optionally set headers (if needed)
			req.Header.Set("Content-Type", "application/json") // or any other content type

			// Create an HTTP client and set a timeout if desired
			client := http.Client{
				Timeout: 10 * time.Second, // Set timeout to 10 seconds
			}

			// Send the request
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer resp.Body.Close() // Ensure the response body is closed

			// Check the response status
			fmt.Println("Response status:", resp.Status)
			*/
			/**

			jsonBody := []byte(`{"client_message": "hello, server!"}`)
			bodyReader := bytes.NewReader(jsonBody)

			requestURL := fmt.Sprintf("http://localhost:%d?id=1234", serverPort)
			req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)

			*/
			// fmt.Println("COUNTER LINK", counterLink)
		}

		// for _, gaugeLink := range gaugesLinks {
		// 	fmt.Println("GAUGE LINK", gaugeLink)
		// }
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
