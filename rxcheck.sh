#!/bin/bash

export $(grep -v '^#' .env | xargs)

case "$1" in
    run)
        go run ./cmd/server/
        ;;
    seed)
        go run ./cmd/seed/main.go
        ;;
    tailwind)
        tailwindcss -i ./static/input.css -o ./static/styles.css
        ;;
    up)
        cd sql/schema; goose turso ${DB_URL} up
        ;;
    down)
        cd sql/schema; goose turso ${DB_URL} down
        ;;
    watch)
        templ generate --watch --cmd="go run ./cmd/server/" --proxy="http://localhost:8080"
        ;;
     *)
        echo "Usage: $0 {run|seed|tailwind|up|down}"
        exit 1
esac
