package main

import(
	_ "github.com/lib/pq"
	"net/http"
	"fmt"
	"os"
	"rxcheck/internal/database"
	"database/sql"

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

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT not found")
		return
	}

	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		return
	}

	cfg := config{
		db: database.New(dbConnection),
	}

	fs := http.FileServer(http.Dir("./static"))
	
	mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	mux.HandleFunc("/", handleApp)

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