package middleware

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

var defaultGzipTypesMap = map[string]struct{}{}
var defaultGzipTypes = []string{
	"text/plain",
	"text/css",
	"text/xml",
	"application/json",
	"image/svg+xml",
	"application/xml",
	"application/xml+rss",
	"text/javascript",
	"application/x-javascript",
	"application/javascript",
}

type gzResponseWriter struct {
	buf        *bytes.Buffer
	bufSize    int
	minSize    int
	header     http.Header
	statusCode int
}

func (gzrw *gzResponseWriter) Write(b []byte) (int, error) {
	gzrw.WriteHeader(http.StatusOK)

	size, err := gzrw.buf.Write(b)
	gzrw.bufSize += size
	return size, err
}

func (gzrw *gzResponseWriter) WriteHeader(statusCode int) {
	gzrw.statusCode = statusCode
}

func (gzrw *gzResponseWriter) Header() http.Header {
	return gzrw.header
}

func (gzrw *gzResponseWriter) isCompressibleContent() bool {
	_, ok := defaultGzipTypesMap[gzrw.header.Get("Content-Type")]
	return ok && gzrw.bufSize > gzrw.minSize
}

func (gzrw *gzResponseWriter) exportHeaderTo(w http.ResponseWriter) {
	for key, values := range gzrw.Header() {
		for _, v := range values {
			w.Header().Set(key, v)
		}
	}
}

func isAcceptGzipEncoding(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

func GzipCompress(minSize int) func(next http.Handler) http.Handler {
	for _, t := range defaultGzipTypes {
		defaultGzipTypesMap[t] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isAcceptGzipEncoding(r) {
				next.ServeHTTP(w, r)
				return
			}

			gzrw := &gzResponseWriter{
				minSize: minSize,
				buf:     new(bytes.Buffer),
				header:  make(http.Header),
			}

			next.ServeHTTP(gzrw, r)

			gzrw.exportHeaderTo(w)

			if gzrw.isCompressibleContent() {
				gzw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				defer gzw.Close()

				w.Header().Set("Content-Encoding", "gzip")
				w.WriteHeader(gzrw.statusCode)
				io.Copy(gzw, gzrw.buf)

				return
			}

			w.WriteHeader(gzrw.statusCode)
			io.Copy(w, gzrw.buf)
		})
	}
}
