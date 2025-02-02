package main

import(
	"net/http"
	"rxcheck/templates"
)

func handleApp(w http.ResponseWriter, r *http.Request) {
	templates.Index("this is a test").Render(r.Context(), w)
}

func handleSwapLogin(w http.ResponseWriter, r *http.Request){
	templates.Login("Login", "/swap_create").Render(r.Context(), w)
}

func handleSwapCreateAccount(w http.ResponseWriter, r *http.Request){
	templates.Login("Create Account", "swao_login").Render(r.Context(), w)
}