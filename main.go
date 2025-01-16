package main

import(
	"net/http"
)

type config struct{

}

func main(){
	const port = "8080"
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	srv.ListenAndServe()
}