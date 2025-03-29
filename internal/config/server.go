package config

import (
	"flag"

	"github.com/caarlos0/env"
)

// Server configuration struct contract.
type ServerConfigs struct {
	APIHost string `env:"ADDRESS"`
}

// Default server configuration values.
var defaultServerConfigs = ServerConfigs{
	APIHost: "localhost:8080",
}

// Returns configuration data for server application.
func ParseServerConfigs() ServerConfigs {
	envConfigs := parseServerEnvData()
	argsConfigs := parseServerArgsConfigs()
	serverConfigs := mergeServerConfigs(envConfigs, argsConfigs)

	if !isLinkCorrect(serverConfigs.APIHost) {
		serverConfigs.APIHost = defaultServerConfigs.APIHost
	}

	return serverConfigs
}

// Returns parsed server configuration data
// from environmental variables.
func parseServerEnvData() ServerConfigs {
	var configs ServerConfigs
	err := env.Parse(&configs)
	if err != nil {
		return ServerConfigs{}
	}
	return configs
}

// Returns parsed server configuration data
// from command line arguments or default
// values if arguments were not provided.
func parseServerArgsConfigs() ServerConfigs {
	apiHostPointer := flag.String("a", defaultServerConfigs.APIHost, "API host")
	flag.Parse()

	return ServerConfigs{
		APIHost: *apiHostPointer,
	}
}

// Merges server env configs and server
// arg configs with prior to env configs.
func mergeServerConfigs(envConfigs ServerConfigs, argsConfigs ServerConfigs) ServerConfigs {
	return ServerConfigs{
		APIHost: selectExistingString(envConfigs.APIHost, argsConfigs.APIHost),
	}
}
