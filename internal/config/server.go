package config

import (
	"flag"
	"os"

	"github.com/caarlos0/env"
)

type ServerConfigs struct {
	APIHost                  string `env:"ADDRESS"`
	StoreInterval            uint   `env:"STORE_INTERVAL"`
	FileStoragePath          string `env:"FILE_STORAGE_PATH"`
	Restore                  bool   `env:"RESTORE"`
	DatabaseConnectionString string `env:"DATABASE_DSN"`
}

var defaultServerConfigs = ServerConfigs{
	APIHost:                  "localhost:8080",
	StoreInterval:            0,
	FileStoragePath:          "metrics.json",
	Restore:                  true,
	DatabaseConnectionString: "",
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
	storeIntervalPointer := flag.Uint("i", defaultServerConfigs.StoreInterval, "File storage interval")
	fileStoragePathPointer := flag.String("f", defaultServerConfigs.FileStoragePath, "File storage path")
	restorePointer := flag.Bool("r", defaultServerConfigs.Restore, "Flag to restore data from file")
	dbConnectionStringPointer := flag.String("d", defaultServerConfigs.DatabaseConnectionString, "Database connection string")

	flag.Parse()

	return ServerConfigs{
		APIHost:                  *apiHostPointer,
		StoreInterval:            *storeIntervalPointer,
		FileStoragePath:          *fileStoragePathPointer,
		Restore:                  *restorePointer,
		DatabaseConnectionString: *dbConnectionStringPointer,
	}
}

// mergeServerConfigs merges server env configs and server
// arg configs with prior to env configs.
func mergeServerConfigs(envConfigs ServerConfigs, argsConfigs ServerConfigs) ServerConfigs {
	restore := func() bool {
		if os.Getenv("RESTORE") == "" {
			return argsConfigs.Restore
		}
		return envConfigs.Restore
	}
	return ServerConfigs{
		APIHost:                  selectExistingString(envConfigs.APIHost, argsConfigs.APIHost),
		StoreInterval:            selectExistingUInt(envConfigs.StoreInterval, argsConfigs.StoreInterval),
		FileStoragePath:          selectExistingString(envConfigs.FileStoragePath, argsConfigs.FileStoragePath),
		DatabaseConnectionString: selectExistingString(envConfigs.DatabaseConnectionString, argsConfigs.DatabaseConnectionString),
		Restore:                  restore(),
	}
}
