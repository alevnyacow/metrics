package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type ServerConfigs struct {
	APIHost string `env:"ADDRESS"`
}

var defaultServerConfigs = ServerConfigs{
	APIHost: "localhost:8080",
}

// parseServerEnvData returns parsed server configuration data
// from environmental variables.
func parseServerEnvData() ServerConfigs {
	var configs ServerConfigs
	err := env.Parse(&configs)
	if err != nil {
		return ServerConfigs{}
	}
	return configs
}

// parseServerArgsConfigs returns parsed server configuration data
// from command line arguments or default
// values if arguments were not provided.
func parseServerArgsConfigs() ServerConfigs {
	apiHostPointer := flag.String("a", defaultServerConfigs.APIHost, "API host")
	flag.Parse()

	return ServerConfigs{
		APIHost: *apiHostPointer,
	}
}

// mergeServerConfigs merges server env configs and server
// arg configs with prior to env configs.
func mergeServerConfigs(envConfigs ServerConfigs, argsConfigs ServerConfigs) ServerConfigs {
	return ServerConfigs{
		APIHost: selectExistingString(envConfigs.APIHost, argsConfigs.APIHost),
	}
}
