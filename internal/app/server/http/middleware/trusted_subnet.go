package middleware

import (
	"log"
	"net"
	"net/http"
)

// TrustedSubnet checks the user's real IP, whether it belongs to a trusted subnet
func TrustedSubnet(subnet string) func(next http.Handler) http.Handler {
	var trustedSubnet *net.IPNet
	var err error

	_, trustedSubnet, err = net.ParseCIDR(subnet)
	if err != nil {
		log.Fatal(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if trustedSubnet != nil {
				realIP := r.Header.Get("X-Real-IP")
				if !trustedSubnet.Contains(net.ParseIP(realIP)) {
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
