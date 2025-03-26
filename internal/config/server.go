package config

import (
	"flag"
)

const defaultAPIHost = "localhost:8080"

func ForServer() (apiHost string) {
	apiHost = *flag.String("a", defaultAPIHost, "API host")
	flag.Parse()
	return
}
