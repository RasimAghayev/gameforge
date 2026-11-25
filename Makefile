.PHONY: build test verify up down run lint format shell migrate seed clean security

# Compose project name keeps parallel containers isolated.
COMPOSE_PROJECT_NAME ?= gameforge-p0

build:
	docker build --target build -t ghcr.io/rasimaghayev/gameforge:build .

test:
	docker build --target test -t ghcr.io/rasimaghayev/gameforge:test .

verify: build test

up: verify
	COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) docker compose -f infrastructure/docker/compose.dev.yaml up -d

down:
	COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) docker compose -f infrastructure/docker/compose.dev.yaml down

run:
	docker run --rm -p 8080:8080 --env-file .env ghcr.io/rasimaghayev/gameforge:runtime

lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v1.60 golangci-lint run ./...

format:
	docker run --rm -v $$(pwd):/app -w /app golang:1.22 gofmt -w .

shell:
	COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) docker compose -f infrastructure/docker/compose.dev.yaml run --server gameforge sh

migrate:
	COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) docker compose -f infrastructure/docker/compose.dev.yaml exec -T gameforge /server migrate

seed:
	COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) docker compose -f infrastructure/docker/compose.dev.yaml exec -T gameforge /server seed

security:
	docker run --rm -v $$(pwd):/app -w /app aquasec/trivy:latest filesystem --exit-code 0 --skip-dirs /app/.git /app

clean:
	@echo "Refusing to delete data volumes. Use 'make down -v' explicitly to remove volumes."
