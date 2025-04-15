package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
)

func New() http.Handler {
	target := "http://goprox-backend:9000" // cible le service Docker "backend"
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid proxy target URL: %s", target)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Proxying request: %s %s", r.Method, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})
}
