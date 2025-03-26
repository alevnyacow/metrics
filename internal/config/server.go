package config

import (
	"flag"
	"net/url"
)

const defaultAPIHost = "localhost:8080"

func ForServer() (apiHost string) {
	apiHost = *flag.String("a", defaultAPIHost, "API host")
	flag.Parse()

	_, err := url.ParseRequestURI(apiHost)
	if err != nil {
		apiHost = defaultAPIHost
	}

	return
}
