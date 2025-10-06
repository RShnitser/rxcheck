package main

import(
	//_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "github.com/tursodatabase/go-libsql"
	"net/http"
	"fmt"
	"os"
	"rxcheck/internal/database"
	"rxcheck/static"
	"database/sql"

	"github.com/joho/godotenv"
)

type config struct{
	db *database.Queries
	jwtSecret string
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT_SECRET must be set")
		return
	}

	dbConnection, err := sql.Open("libsql", dbURL)
	if err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		return
	}

	cfg := config{
		db: database.New(dbConnection),
		jwtSecret: jwtSecret,
	}

	//fs := http.FileServer(http.Dir("./static"))
	fs := http.FileServer(http.FS(static.StaticFiles))
	
	mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	mux.HandleFunc("/", handleApp)
	mux.HandleFunc("/swap_login", handleSwapLogin)
	mux.HandleFunc("/swap_create", handleSwapCreateAccount)
	mux.HandleFunc("/create_user", cfg.handleCreateUser)
	mux.HandleFunc("/login_user", cfg.handleLogin)
	mux.HandleFunc("/quiz/{drugClassification}", cfg.handleCreateQuiz)
	mux.HandleFunc("/question", cfg.handleGetQuestion)
	mux.HandleFunc("/menu", cfg.handleGetMenu)

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