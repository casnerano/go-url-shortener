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

// Write to buffer.
func (gzrw *gzResponseWriter) Write(b []byte) (int, error) {
	if gzrw.statusCode == 0 {
		gzrw.WriteHeader(http.StatusOK)
	}

	size, err := gzrw.buf.Write(b)
	gzrw.bufSize += size
	return size, err
}

// WriteHeader - set status code.
func (gzrw *gzResponseWriter) WriteHeader(statusCode int) {
	gzrw.statusCode = statusCode
}

// Header getter.
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

// GzipCompress middleware for gzip compressing the response body.
//
// Applies only if the following conditions are met:
//  1. The `Accept-Encoding` request header contains the `gzip` value.
//  2. The `Content-Type` response header contains one of the allowed values. See `defaultGzipTypes` constant.
//  3. The size of the raw response body is greater than the value set in the `minSize` argument.
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
