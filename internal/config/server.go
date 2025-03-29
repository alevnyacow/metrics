package config

import (
	"flag"

	"github.com/caarlos0/env"
)

const defaultAPIHost = "localhost:8080"

type ServerEnvConfigs struct {
	ApiHost string `env:"ADDRESS"`
}

func parseServerEnvData() (apiHost string) {
	var configs ServerEnvConfigs
	err := env.Parse(&configs)
	if err != nil {
		return
	}
	apiHost = configs.ApiHost
	return
}

func ForServer() (apiHost string) {
	envDataApiHost := parseServerEnvData()
	apiHostPointer := flag.String("a", defaultAPIHost, "API host")
	flag.Parse()

	apiHost = selectExistingString(*apiHostPointer, envDataApiHost)
	isCorrectLink := checkLink(apiHost)

	if !isCorrectLink {
		apiHost = defaultAPIHost
	}

	return
}
