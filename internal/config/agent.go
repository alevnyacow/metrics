package config

import (
	"flag"
	"strings"

	"github.com/caarlos0/env"
)

const defaultClientAPIHost = "localhost:8080"
const defaultPollInterval = 10
const defaultReportInterval = 2

type AgentEnvConfigs struct {
	APIHost        string `env:"ADDRESS"`
	ReportInterval uint   `env:"REPORT_INTERVAL"`
	PollInterval   uint   `env:"POLL_INTERVAL"`
}

// Returns parsed agent configuration data
// from environmental variables.
func parseAgentEnvData() (apiHost string, pollInterval uint, reportInterval uint) {
	var envConfigs AgentEnvConfigs
	err := env.Parse(&envConfigs)
	if err != nil {
		return
	}
	apiHost = envConfigs.APIHost
	pollInterval = envConfigs.PollInterval
	reportInterval = envConfigs.ReportInterval
	return
}

// Returns parsed agent configuration data
// from command line arguments or default
// values if arguments were not provided.
func parseAgentParams() (apiHost string, pollInterval uint, reportInterval uint) {
	apiHostPointer := flag.String("a", defaultClientAPIHost, "API host")
	pollIntervalPointer := flag.Uint("p", defaultPollInterval, "Poll interval")
	reportIntervalPointer := flag.Uint("r", defaultReportInterval, "Report interval")
	flag.Parse()

	apiHost = *apiHostPointer
	pollInterval = *pollIntervalPointer
	reportInterval = *reportIntervalPointer

	return
}

// Returns run configuration data for agent application.
func ForAgent() (apiHost string, pollInterval uint, reportInterval uint) {
	apiHostFromEnv, pollIntervalFromEnv, reportIntervalFromEnv := parseAgentEnvData()
	apiHostFromParams, pollIntervalFromParams, reportIntervalFromParams := parseAgentParams()

	apiHost = selectExistingString(apiHostFromEnv, apiHostFromParams)
	pollInterval = selectExistingUInt(pollIntervalFromEnv, pollIntervalFromParams)
	reportInterval = selectExistingUInt(reportIntervalFromEnv, reportIntervalFromParams)

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
