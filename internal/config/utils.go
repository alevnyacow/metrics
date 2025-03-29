package config

import (
	"net/url"
	"strings"
)

// Checks if link is correct (trying to parse as URI)
// and checks if link starts with http prefix, which
// covers both "http://" and "https://".
func isLinkCorrect(target string) (isCorrect bool) {
	_, err := url.ParseRequestURI(target)
	isCorrect = err == nil
	return
}

// Checks if provided link string is localhost
// and it does not start with "http://".
func isLocalhostWithoutPrefix(target string) bool {
	return strings.HasPrefix(target, "localhost:")
}

// Returns given link with added "http://" prefix.
func withHTTPPrefix(target string) string {
	return "http://" + target
}

// If first string is empty, second string is returned.
// First string is returned otherwise.
func selectExistingString(lhs string, rhs string) string {
	if lhs == "" {
		return rhs
	}

	return lhs
}

// If first uint is zero, second string is returned.
// First string is returned otherwise.
func selectExistingUInt(lhs uint, rhs uint) uint {
	if lhs == 0 {
		return rhs
	}

	return lhs
}
