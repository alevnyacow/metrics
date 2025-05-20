package middleware

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/alevnyacow/metrics/internal/hash"
	"github.com/rs/zerolog/log"
)

func WithHashCheck(key string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		if key == "" {
			return handler
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hashData := r.Header.Get("HashSHA256")
			if hashData != "" {
				body, bodyReadingError := io.ReadAll(r.Body)
				if bodyReadingError != nil {
					log.Err(bodyReadingError).Msg("Error on body reading")
					return
				}
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				hashedBody, hashError := hash.SignedSHA256(body, key)
				if hashError != nil {
					log.Err(hashError).Msg("Hash error")
				}
				w.Header().Add("HashSHA256", hashedBody)
				if hashedBody != hashData {
					log.Err(errors.New("hashes are not the same")).Msg("Access denied - hashes are not the same")
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			handler.ServeHTTP(w, r)
		})
	}
}
