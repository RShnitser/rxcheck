# RXCHECK

RxCheck is a quiz application focusing on multiple choice medication related questions
intended to help prepare for the PTCB exam.

## Installation

Requires Go

[Installing Go](https://go.dev/doc/install)

Requires goose, sqlc, and templ

```
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/a-h/templ/cmd/templ@latest
```
create .env file with these variables:
PORT : port number
DB_URL : turso database url
JWT_SECRET : randomly generated key

## Running

Run Database Migration:
```
./rxcheck.sh up
```

Run Database Seed:
```
./rxcheck.sh seed
```

Run Server:
```
./rxcheck.sh run
```