package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cmehx/goprox/internal/proxy"
)

func main() {
	target := os.Getenv("PROXY_TARGET")
	if target == "" {
		log.Fatal("PROXY_TARGET not set")
	}

	handler := proxy.New(target)

	log.Println("Starting proxy on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
