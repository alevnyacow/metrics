package api

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type compressWriter struct {
	writer     http.ResponseWriter
	gzipWriter *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		writer:     w,
		gzipWriter: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.writer.Header()
}

func isCompatibleContentType(contentTypeHeaderContent string) bool {
	supportedContentTypes := [2]string{"application/json", "text/html"}
	for _, contentType := range supportedContentTypes {
		if strings.Contains(contentTypeHeaderContent, contentType) {
			return true
		}
	}
	return false
}

func (c *compressWriter) Write(p []byte) (int, error) {
	if isCompatibleContentType(c.Header().Get("Content-Type")) {
		return c.gzipWriter.Write(p)
	}
	return c.writer.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 && isCompatibleContentType(c.Header().Get("Content-Type")) {
		c.writer.Header().Set("Content-Encoding", "gzip")
	}
	c.writer.WriteHeader(statusCode)
}

func (c *compressWriter) Close() error {
	return c.gzipWriter.Close()
}

func withGZipResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ow := w
		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			cw := newCompressWriter(w)
			ow = cw
			defer cw.Close()
		}
		handler.ServeHTTP(ow, r)
	})
}
