.PHONY: up down rebuild reset-db logs logs-api db \
	build run-local test test-race lint fmt tidy \
	migrate-up migrate-down create-migration

# -------------------------
# Docker
# -------------------------

up:
	docker compose up

down:
	docker compose down

rebuild:
	docker compose up --build

reset-db:
	docker compose down -v

logs:
	docker compose logs -f

logs-api:
	docker compose logs -f api

db:
	docker compose exec db psql -U postgres -d postgres


# -------------------------
# Go development
# -------------------------

build:
	go build -o bin/server ./cmd/server

run-local:
	go run ./cmd/server

test:
	go test ./...

test-race:
	go test -race ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

tidy:
	go mod tidy


# -------------------------
# Migrations
# -------------------------

migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down 1

create-migration:
ifndef name
	$(error name is required. Example: make create-migration name=create_players_table)
endif
	migrate create -ext sql -dir migrations -seq $(name)