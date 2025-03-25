package api

import (
	"net/http"
)

func processUpdatePathParsingResult(parsingResult updatePathParsingResult, responseWriter http.ResponseWriter) (finishedWithError bool) {
	if !parsingResult.parsedName {
		finishedWithError = true
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}

	if !parsingResult.parsedValue {
		finishedWithError = true
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	finishedWithError = false
	return
}
