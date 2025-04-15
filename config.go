package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contient la configuration de l'application
type Config struct {
	ProxyTarget string
	Port        string
}

// Load charge la configuration depuis le fichier .env
func Load() Config {
	// Charger les variables d'environnement depuis le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Charger les variables d'environnement
	target := os.Getenv("PROXY_TARGET")
	if target == "" {
		log.Fatal("Missing PROXY_TARGET environment variable")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8180" // Port par défaut si non spécifié
	}

	// Retourner la configuration
	return Config{
		ProxyTarget: target,
		Port:        port,
	}
}
