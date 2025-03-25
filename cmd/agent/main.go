package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alevnyacow/metrics/internal/generator"
)

func generateMetrics(intervalInSeconds int, counters *generator.Counters, gauges *generator.Gauges) {
	for {
		time.Sleep(time.Duration(intervalInSeconds) * time.Second)
		*counters = generator.GenerateCounters()
		*gauges = generator.GenerateGauges()
	}
}

func sendMetrics(intervalInSeconds int, apiRoot string, counters generator.WithLinks, gauges generator.WithLinks) {
	for {
		time.Sleep(time.Duration(intervalInSeconds) * time.Second)

		countersLinks := counters.Links(apiRoot)
		// gaugesLinks := gauges.Links(apiRoot)

		for _, counterLink := range countersLinks {
			fmt.Println("CALLED " + counterLink)
			req, g := http.NewRequest("POST", counterLink, nil)
			if g != nil {
				fmt.Println(g.Error())
			}
			client := http.Client{
				Timeout: 10 * time.Second, // Set timeout to 10 seconds
			}
			client.Do(req)

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

	pollInterval := 2
	reportInterval := 10

	go generateMetrics(pollInterval, counterMetrics, gaugeMetrics)
	go sendMetrics(reportInterval, "http://localhost:8080", counterMetrics, gaugeMetrics)

	select {}
}
