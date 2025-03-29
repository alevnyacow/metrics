package config

import (
	"flag"

	"github.com/caarlos0/env"
)

// Agent configuration struct contract.
type AgentConfigs struct {
	APIHost        string `env:"ADDRESS"`
	ReportInterval uint   `env:"REPORT_INTERVAL"`
	PollInterval   uint   `env:"POLL_INTERVAL"`
}

// Default client configuration values.
var defaultAgentConfigs = AgentConfigs{
	APIHost:        "localhost:8080",
	ReportInterval: 2,
	PollInterval:   10,
}

// Returns configuration data for agent application.
func ParseAgentConfigs() AgentConfigs {
	envConfigs := parseAgentEnvData()
	argsConfigs := parseAgentArgsConfigs()
	agentConfigs := mergeClientConfigs(envConfigs, argsConfigs)

	if !isLinkCorrect(agentConfigs.APIHost) {
		agentConfigs.APIHost = defaultAgentConfigs.APIHost
		return agentConfigs
	}

	if isLocalhostWithoutPrefix(agentConfigs.APIHost) {
		agentConfigs.APIHost = withHTTPPrefix(agentConfigs.APIHost)
	}

	return agentConfigs
}

// Returns parsed agent configuration data
// from environmental variables.
func parseAgentEnvData() AgentConfigs {
	var envConfigs AgentConfigs
	err := env.Parse(&envConfigs)
	if err != nil {
		return AgentConfigs{}
	}
	return envConfigs
}

// Returns parsed agent configuration data
// from command line arguments or default
// values if arguments were not provided.
func parseAgentArgsConfigs() AgentConfigs {
	apiHostPointer := flag.String("a", defaultAgentConfigs.APIHost, "API host")
	pollIntervalPointer := flag.Uint("p", defaultAgentConfigs.PollInterval, "Poll interval")
	reportIntervalPointer := flag.Uint("r", defaultAgentConfigs.ReportInterval, "Report interval")
	flag.Parse()

	return AgentConfigs{
		APIHost:        *apiHostPointer,
		PollInterval:   *pollIntervalPointer,
		ReportInterval: *reportIntervalPointer,
	}
}

// Merges agent env configs and agent
// arg configs with prior to env configs.
func mergeClientConfigs(envConfigs AgentConfigs, argsConfigs AgentConfigs) AgentConfigs {
	return AgentConfigs{
		APIHost:        selectExistingString(envConfigs.APIHost, argsConfigs.APIHost),
		ReportInterval: selectExistingUInt(envConfigs.ReportInterval, argsConfigs.ReportInterval),
		PollInterval:   selectExistingUInt(envConfigs.PollInterval, argsConfigs.PollInterval),
	}
}
