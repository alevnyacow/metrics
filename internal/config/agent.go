package config

import (
	"flag"
	"net/url"
)

const defaultClientAPIHost = "http://localhost:8080"
const defaultPollInterval = 10
const defaultReportInterval = 2

func ForAgent() (apiHost string, pollInterval uint, reportInterval uint) {
	apiHost = *flag.String("a", defaultClientAPIHost, "API host")
	pollInterval = *flag.Uint("p", defaultPollInterval, "Poll interval")
	reportInterval = *flag.Uint("r", defaultReportInterval, "Report interval")
	flag.Parse()

	_, err := url.ParseRequestURI(apiHost)
	if err != nil {
		apiHost = defaultAPIHost
	}

	return
}
