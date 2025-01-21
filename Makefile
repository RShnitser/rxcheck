include .env

.PHONY: up
up:
	cd sql/schema; goose postgres ${DB_URL} up

.PHONY: down
down:
	cd sql/schema; goose postgres ${DB_URL} down