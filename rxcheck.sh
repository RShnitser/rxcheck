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
        tailwindcss -i ./static/input.css -o ./static/output.css
        ;;
    up)
        cd sql/schema; goose postgres ${DB_URL} up
        ;;
    down)
        cd sql/schema; goose postgres ${DB_URL} down
        ;;
     *)
        echo "Usage: $0 {run|seed|tailwind|up|down}"
        exit 1
esac
