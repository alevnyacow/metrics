package config

import (
	"flag"
)

const defaultAPIHost = "localhost:8080"

func ForServer() (apiHost string) {
	apiHostPointer := flag.String("a", defaultAPIHost, "API host")
	flag.Parse()

	apiHost = *apiHostPointer
	isCorrectLink := checkLink(apiHost)

	if !isCorrectLink {
		apiHost = defaultAPIHost
	}

	return
}
