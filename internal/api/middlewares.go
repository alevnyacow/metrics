package api

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type responseData struct {
	contentLength int
	statusCode    int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (loggingW *loggingResponseWriter) Write(data []byte) (int, error) {
	size, error := loggingW.ResponseWriter.Write(data)
	loggingW.responseData.contentLength += size
	return size, error
}

func (loggingW *loggingResponseWriter) WriteHeader(statusCode int) {
	loggingW.ResponseWriter.WriteHeader(statusCode)
	loggingW.responseData.statusCode = statusCode
}

func withLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseData := &responseData{contentLength: 0, statusCode: 200}
		loggingW := loggingResponseWriter{ResponseWriter: w, responseData: responseData}
		handler.ServeHTTP(&loggingW, r)
		duration := time.Since(start)
		log.Info().Str("URI", r.RequestURI).Str("Method", r.Method).Dur("Duration", duration).Msg("Request")
		log.Info().Int("Status", responseData.statusCode).Int("Content length", responseData.contentLength).Msg("Response")
	})
}
