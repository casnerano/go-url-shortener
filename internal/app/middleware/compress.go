package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

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
	http.ResponseWriter
	Writer io.Writer
}

func (gzw gzResponseWriter) Write(b []byte) (int, error) {
	return gzw.Writer.Write(b)
}

func GzipCompress() func(next http.Handler) http.Handler {

	allowedGzipTypes := make(map[string]struct{})
	for _, t := range defaultGzipTypes {
		allowedGzipTypes[t] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				if _, ok := allowedGzipTypes[r.Header.Get("Content-Type")]; ok {
					gzw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
					defer gzw.Close()

					w.Header().Set("Content-Encoding", "gzip")
					w = gzResponseWriter{ResponseWriter: w, Writer: gzw}
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
