package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	contentLength int
	statusCode    int
}

func (loggingW *loggingResponseWriter) Write(data []byte) (int, error) {
	size, error := loggingW.ResponseWriter.Write(data)
	loggingW.contentLength += size
	return size, error
}

func (loggingW *loggingResponseWriter) WriteHeader(statusCode int) {
	loggingW.ResponseWriter.WriteHeader(statusCode)
	loggingW.statusCode = statusCode
}

func WithLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		loggingW := loggingResponseWriter{ResponseWriter: w, contentLength: 0, statusCode: 200}
		handler.ServeHTTP(&loggingW, r)
		duration := time.Since(start)
		log.Info().Str("URI", r.RequestURI).Str("Method", r.Method).Dur("Duration", duration).Msg("Request")
		log.Info().Int("Status", loggingW.statusCode).Int("Content length", loggingW.contentLength).Msg("Response")
	})
}
