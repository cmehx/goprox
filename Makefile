# Définir les variables pour faciliter la gestion des commandes

# Nom du service Go
GO_SERVICE = goprox
# Le dossier contenant le Dockerfile
DOCKERFILE_PATH = .
# Nom du conteneur Docker
CONTAINER_NAME = goprox

# Commandes par défaut pour Docker
DOCKER_BUILD = docker compose build
DOCKER_UP = docker compose up
DOCKER_DOWN = docker compose down
DOCKER_RUN = docker compose run --rm $(GO_SERVICE)

# Commande pour construire l'image Docker et lancer le conteneur
build:
	$(DOCKER_BUILD)

# Commande pour démarrer les services (incluant le conteneur Go avec Air)
up:
	$(DOCKER_UP)

# Commande pour stopper les services Docker
down:
	$(DOCKER_DOWN)

# Commande pour démarrer seulement le service Go avec Air
dev:
	$(DOCKER_RUN) air -c .air.toml

# Commande pour construire et tester le projet dans Docker
test:
	$(DOCKER_RUN) go test ./...

# Commande pour démarrer le conteneur en mode développement avec Air
docker-dev:
	$(DOCKER_RUN) air -c .air.toml

# Commande pour afficher le statut des services Docker
status:
	docker ps

# Commande pour démarrer un shell interactif dans le conteneur Go
shell:
	$(DOCKER_RUN) sh

# Commande pour supprimer les volumes Docker (utile pour réinitialiser l'état)
clean:
	docker-compose down -v
