package api

import "net/http"

// marshalingErrorResponse takes error as a parameter
// and returns handler function which returns 500 status
// code and error body.
func marshalingErrorResponse(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// unknownMetricTypeResponse returns handler function
// in case server could not recognize metric type
// provided by user.
func unknownMetricTypeResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// nonExistingMetricOfKnownTypeResponse returns handler function
// in case server recognized metric type provided by user,
// but user requested non-existing metric of this type.
func nonExistingMetricOfKnownTypeResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
}

// notProvidedUpdateValueResponse returns handler function
// in case user did not provide metric value when requested
// update of this metric.
func notProvidedUpdateValueResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// providedIncorrectUpdateValueResponse returns handler function
// in case user provided raw string metric value when requested
// update of this metric but server could not parse it to
// actual metric value.
func providedIncorrectUpdateValueResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
}
