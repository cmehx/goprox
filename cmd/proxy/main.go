package main

import (
	"os"
	"log"
	"net/http"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/cmehx/goprox/internal/proxy"
)

func main() {
	// Charger la configuration
	cfg := godotenv.Load()
	if cfg != nil {
		log.Fatal("Error loading .env file")
	}
	// Afficher un message de démarrage
	log.Printf("Starting proxy on :%s -> %s\n", os.Getenv("PORT"), os.Getenv("PROXY_TARGET"))

	port := os.Getenv("PORT")
	target := os.Getenv("PROXY_TARGET")

	log.Printf("Starting proxy on :%s -> %s\n", port, target)

	// Handler par défaut
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from backend")
	})

	// Reverse proxy en fallback sur "/"
	http.Handle("/", proxy.New())

	// Ici on utilise le mux par défaut
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
