package config

import (
	"net/url"
	"strings"
)

// isLinkCorrect tries to parse target as URI to
// verify if target is correct
func isLinkCorrect(target string) (isCorrect bool) {
	_, err := url.ParseRequestURI(target)
	isCorrect = err == nil
	return
}

// isLocalhostWithoutPrefix checks if provided link string
// is localhost and it does not start with "http://".
func isLocalhostWithoutPrefix(target string) bool {
	return strings.HasPrefix(target, "localhost:")
}

// withHTTPPrefix returns given link with added "http://" prefix.
func withHTTPPrefix(target string) string {
	return "http://" + target
}

// selectExistingString returns second string if first string is empty.
// First string is returned otherwise.
func selectExistingString(lhs string, rhs string) string {
	if lhs == "" {
		return rhs
	}

	return lhs
}

// selectExistingUInt returns second uint if first uint is zero.
// First uint is returned otherwise.
func selectExistingUInt(lhs uint, rhs uint) uint {
	if lhs == 0 {
		return rhs
	}

	return lhs
}
