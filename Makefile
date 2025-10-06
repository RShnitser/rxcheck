include .env

.PHONY: up
up:
	cd sql/schema; goose turso ${DB_URL} up

.PHONY: down
down:
	cd sql/schema; goose turso ${DB_URL} down

.PHONY: tailwind
tailwind:
	tailwindcss -i input.css -o ./static/styles.css

.PHONY: run
run:
	go run -tags=local ./cmd/server/ 

.PHONY: seed
seed:
	go run ./cmd/seed/main.go

.PHONY: gbuild
gbuild:
	go build ./cmd/server/

.PHONY: dbuild
dbuild:
	docker build . -t rxcheck:latest

.PHONY: drun
drun:
	docker run --env-file=.env -p 8080:8080 rxcheck