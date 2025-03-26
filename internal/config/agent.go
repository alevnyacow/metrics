package config

import (
	"flag"
	"net/url"
)

const defaultClientAPIHost = "http://localhost:8080"
const defaultPollInterval = 10
const defaultReportInterval = 2

func ForAgent() (apiHost string, pollInterval uint, reportInterval uint) {
	apiHostPointer := flag.String("a", defaultClientAPIHost, "API host")
	pollIntervalPointer := flag.Uint("p", defaultPollInterval, "Poll interval")
	reportIntervalPointer := flag.Uint("r", defaultReportInterval, "Report interval")
	flag.Parse()

	apiHost = *apiHostPointer
	pollInterval = *pollIntervalPointer
	reportInterval = *reportIntervalPointer

	_, err := url.ParseRequestURI(apiHost)
	if err != nil {
		apiHost = defaultAPIHost
	}

	return
}
