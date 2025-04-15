package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic("Invalid PROXY_TARGET: " + target)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging ou extensions futures
		proxy.ServeHTTP(w, r)
	})
}
