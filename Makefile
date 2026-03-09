.PHONY: run stop rebuild reset-db logs

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

create-migration:
ifndef name
	$(error name is required. Example: make create-migration name=create_players_table)
endif
	migrate create -ext sql -dir migrations -seq $(name)