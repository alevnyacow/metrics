package api

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	gzipW *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.gzipW.Write(b)
}

func handleGzipRequest(w http.ResponseWriter, r *http.Request, handler http.Handler) {
	gzReader, err := gzip.NewReader(r.Body)
	if err != nil {
		unparsebleCompressedDataFromClient()(w, r)
		return
	}
	defer func() {
		err := gzReader.Close()
		if err != nil {
			unparsebleCompressedDataFromClient()(w, r)
		}
	}()

	r.Body = gzReader
	w.Header().Set("Content-Encoding", "gzip")
	gzipWriter := gzip.NewWriter(w)
	defer func() {
		err := gzipWriter.Close()
		if err != nil {
			unparsebleCompressedDataFromClient()(w, r)
		}
	}()

	gzipResponseWriter := gzipResponseWriter{gzipW: gzipWriter, ResponseWriter: w}
	handler.ServeHTTP(gzipResponseWriter, r)
}

func handleGzipResponse(w http.ResponseWriter, r *http.Request, handler http.Handler) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/html")
	gw := gzip.NewWriter(w)
	defer func() {
		err := gw.Close()
		if err != nil {
			compressingDataTrouble()(w, r)
		}
	}()

	gzipWriter := gzipResponseWriter{gzipW: gw, ResponseWriter: w}
	handler.ServeHTTP(gzipWriter, r)
}

func withGzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientSendedGZip := strings.Contains(r.Header.Get("Content-Encoding"), "gzip")
		if clientSendedGZip {
			handleGzipRequest(w, r, handler)
			return
		}

		clientAcceptsGZip := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
		clientAcceptsHTML := strings.Contains(r.Header.Get("Accept"), "text/html")
		returnCompressedData := clientAcceptsGZip && clientAcceptsHTML
		if returnCompressedData {
			handleGzipResponse(w, r, handler)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
