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

func selectExistingString(lhs string, rhs string) string {
	if lhs == "" {
		return rhs
	}

	return lhs
}

func selectExistingUInt(lhs uint, rhs uint) uint {
	if lhs == 0 {
		return rhs
	}

	return lhs
}
