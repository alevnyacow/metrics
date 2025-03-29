package config

import (
	"net/url"
)

// Checks if link is correct (trying to parse as URI)
// and checks if link starts with http prefix, which
// covers both "http://" and "https://".
func checkLink(target string) (isCorrect bool) {
	_, err := url.ParseRequestURI(target)
	isCorrect = err == nil
	return
}
