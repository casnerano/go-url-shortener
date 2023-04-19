package middleware

import (
	"compress/gzip"
	"net/http"
)

// GzipDecompress middleware for gzip decompressing the request body.
// Only applies if there is a `Content-Encoding` header with a value of `gzip`.
func GzipDecompress() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Encoding") == "gzip" {
				gzr, err := gzip.NewReader(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				defer gzr.Close()
				r.Body = gzr
			}
			next.ServeHTTP(w, r)
		})
	}
}
