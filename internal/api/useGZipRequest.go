package api

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type compressReader struct {
	reader     io.ReadCloser
	gzipReader *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		reader:     r,
		gzipReader: gzipReader,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.gzipReader.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.gzipReader.Close(); err != nil {
		return err
	}
	return c.gzipReader.Close()
}

func withGZipRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			contentEncoding := r.Header.Get("Content-Encoding")
			sendsGzip := strings.Contains(contentEncoding, "gzip")
			if sendsGzip {
				cr, err := newCompressReader(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				r.Body = cr
				defer cr.Close()
			}
			handler.ServeHTTP(w, r)
		},
	)
}
