package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	// Charger la configuration
	cfg := config.Load()

	// Afficher un message de démarrage
	log.Printf("Starting proxy on :%s -> %s\n", cfg.Port, cfg.ProxyTarget)

	// Créer le handler pour le proxy
	handler := proxy.New(cfg.ProxyTarget)

	// Lancer le serveur HTTP
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
