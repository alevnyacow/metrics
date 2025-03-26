package config

import (
	"flag"
	"net/url"
)

const defaultAPIHost = "localhost:8080"

func ForServer() (apiHost string) {
	apiHostPointer := flag.String("a", defaultAPIHost, "API host")
	flag.Parse()
	apiHost = *apiHostPointer

	_, err := url.ParseRequestURI(apiHost)
	if err != nil {
		apiHost = defaultAPIHost
	}

	return
}
