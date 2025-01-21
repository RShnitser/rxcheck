package main

import(
	_ "github.com/lib/pq"
	"net/http"
	"fmt"
	"os"
	"rxcheck/templates"
	"rxcheck/internal/database"
	"database/sql"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type config struct{
	db *database.Queries
}

func main(){
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}

	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		return
	}

	const port = "8080"

	cfg := config{
		db: database.New(dbConnection),
	}

	mux := http.NewServeMux()
	
	index := templates.Index("test")
	
	mux.Handle("/", templ.Handler(index))

	mux.HandleFunc("POST /api/users", cfg.handleAddUser)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Unable to start server: %v\n", err)
	}
}