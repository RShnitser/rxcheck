package main

import _ "github.com/lib/pq"
import(
	"net/http"
	"fmt"
	"github.com/a-h/templ"
)

type config struct{

}

func main(){
	const port = "8080"
	mux := http.NewServeMux()
	
	indexComponent := index("test")
	mux.Handle("/", templ.Handler(indexComponent))

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