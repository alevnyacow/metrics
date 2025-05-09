package api

import (
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

// marshalingErrorResponse takes error as a parameter
// and returns handler function for case where server
// could obtain requested data but could not serialize
// this data to JSON for response.
func marshalingErrorResponse(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(err).Msg("Error on marshaling")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// unknownMetricTypeResponse returns handler function
// for case where server could not recognize metric type
// provided by user.
func unknownMetricTypeResponse(unknownMetricType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(errors.New(unknownMetricType)).Msg("Unknown metric type")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// nonExistingMetricOfKnownTypeResponse returns handler function
// for case where server recognized metric type provided by user,
// but user requested non-existing metric of this type.
func nonExistingMetricOfKnownTypeResponse(metricName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(errors.New(metricName)).Msg("Non existing metric request")
		w.WriteHeader(http.StatusNotFound)
	}
}

// notProvidedUpdateValueResponse returns handler function
// for case where user did not provide metric value when requested
// update of this metric.
func notProvidedUpdateValueResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(errors.New("no update value")).Msg("Error on metric update")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// providedIncorrectUpdateValueResponse returns handler function
// for case where user provided raw string metric value when requested
// update of this metric but server could not parse it to
// actual metric value.
func providedIncorrectUpdateValueResponse(notParsedValue string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(errors.New(notParsedValue + " string value could not be parsed")).Msg("Error on metric update")
		w.WriteHeader(http.StatusBadRequest)
	}
}

// failedDatabasePingResponse returns handler function
// for case where client tried to ping database and it
// was not successful
func failedDatabasePingResponse(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(err).Msg("Error on database ping")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func failedUpdatesResponse(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Err(err).Msg("Error on update with batch")
		w.WriteHeader(http.StatusBadRequest)
	}
}
