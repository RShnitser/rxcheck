package main

import(
	_ "github.com/lib/pq"
	"net/http"
	"fmt"
	"os"
	"rxcheck/templates"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type config struct{

}

func main(){
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}

	const port = "8080"
	mux := http.NewServeMux()
	
	index := templates.Index("test")
	mux.Handle("/", templ.Handler(index))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on port %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Unable to start server: %v\n", err)
	}
}