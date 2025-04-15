# Choisir l'image Go officielle, version >= 1.16
FROM golang:1.23-alpine

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Installer Air
RUN go install github.com/air-verse/air@latest
RUN go install github.com/joho/godotenv/cmd/godotenv@latest
# Copier go.mod et go.sum dans le conteneur pour gérer les dépendances
COPY go.mod go.sum ./
# Télécharger les dépendances Go (utilise le cache Docker pour améliorer les performances)
RUN go mod download

# Copier le code source dans le conteneur (cmd, internal, pkg, etc.)
COPY . .

RUN go mod tidy
# Lancer Air avec le fichier de configuration .air.toml (à définir dans ton projet)
CMD ["air", "-c", ".air.toml"]
