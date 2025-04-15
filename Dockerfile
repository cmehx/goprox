# Étape 1 : build l'app
FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o proxy ./cmd/proxy

# Étape 2 : image finale minimale
FROM debian:bookworm-slim

# Crée un user non-root (bonne pratique)
RUN useradd -m appuser

COPY --from=builder /app/proxy /usr/local/bin/proxy

USER appuser
EXPOSE 8080

CMD ["proxy"]
