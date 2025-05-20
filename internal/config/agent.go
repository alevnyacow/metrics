package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type AgentConfigs struct {
	APIHost        string `env:"ADDRESS"`
	ReportInterval uint   `env:"REPORT_INTERVAL"`
	PollInterval   uint   `env:"POLL_INTERVAL"`
	Key            string `env:"KEY"`
}

var defaultAgentConfigs = AgentConfigs{
	APIHost:        "localhost:8080",
	ReportInterval: 10,
	PollInterval:   2,
	Key:            "",
}

// parseAgentEnvData returns parsed agent configuration data
// from environmental variables.
func parseAgentEnvData() AgentConfigs {
	var envConfigs AgentConfigs
	err := env.Parse(&envConfigs)
	if err != nil {
		return AgentConfigs{}
	}
	return envConfigs
}

// parseAgentArgsConfigs returns parsed agent configuration data
// from command line arguments or default
// values if arguments were not provided.
func parseAgentArgsConfigs() AgentConfigs {
	apiHostPointer := flag.String("a", defaultAgentConfigs.APIHost, "API host")
	pollIntervalPointer := flag.Uint("p", defaultAgentConfigs.PollInterval, "Poll interval")
	reportIntervalPointer := flag.Uint("r", defaultAgentConfigs.ReportInterval, "Report interval")
	keyPointer := flag.String("k", defaultAgentConfigs.Key, "SHA Key")
	flag.Parse()

	return AgentConfigs{
		APIHost:        *apiHostPointer,
		PollInterval:   *pollIntervalPointer,
		ReportInterval: *reportIntervalPointer,
		Key:            *keyPointer,
	}
}

// mergeClientConfigs merges agent env configs and agent
// arg configs with prior to env configs.
func mergeClientConfigs(envConfigs AgentConfigs, argsConfigs AgentConfigs) AgentConfigs {
	return AgentConfigs{
		APIHost:        selectExistingString(envConfigs.APIHost, argsConfigs.APIHost),
		ReportInterval: selectExistingUInt(envConfigs.ReportInterval, argsConfigs.ReportInterval),
		PollInterval:   selectExistingUInt(envConfigs.PollInterval, argsConfigs.PollInterval),
		Key:            selectExistingString(envConfigs.Key, argsConfigs.Key),
	}
}
