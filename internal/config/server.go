package config

import (
	"flag"

	"github.com/caarlos0/env"
)

const defaultAPIHost = "localhost:8080"

type ServerEnvConfigs struct {
	APIHost string `env:"ADDRESS"`
}

func parseServerEnvData() (apiHost string) {
	var configs ServerEnvConfigs
	err := env.Parse(&configs)
	if err != nil {
		return
	}
	apiHost = configs.APIHost
	return
}

func ForServer() (apiHost string) {
	envDataAPIHost := parseServerEnvData()
	apiHostPointer := flag.String("a", defaultAPIHost, "API host")
	flag.Parse()

	apiHost = selectExistingString(envDataAPIHost, *apiHostPointer)
	isCorrectLink := checkLink(apiHost)

	if !isCorrectLink {
		apiHost = defaultAPIHost
	}

	return
}
