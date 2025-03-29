package config

import (
	"flag"
	"strings"
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

	isCorrectLink := checkLink(apiHost)

	if !isCorrectLink {
		apiHost = defaultAPIHost
		return
	}

	if isLocalhostWithoutPrefix(apiHost) {
		apiHost = withHTTPPrefix(apiHost)
	}

	return
}

// Checks if provided link string is localhost
// and it does not start with "http://".
func isLocalhostWithoutPrefix(target string) bool {
	return strings.HasPrefix(target, "localhost:")
}

// Returns given link with added "http://" prefix.
func withHTTPPrefix(target string) string {
	return "http://" + target
}
