include .env

.PHONY: up
up:
	cd sql/schema; goose postgres ${DB_URL} up

.PHONY: down
down:
	cd sql/schema; goose postgres ${DB_URL} down

.PHONY: tailwind
tailwind:
	tailwindcss -i ./static/input.css -o ./static/output.css

.PHONY: run
run:
	go run ./cmd/server/