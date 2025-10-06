//go:build !local

package main

import(
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"database/sql"
)

func GetDBConnection(dbURL string)(*sql.DB, error){
	return sql.Open("libsql", dbURL)
}