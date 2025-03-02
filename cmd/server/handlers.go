package main

import(
	"net/http"
	"rxcheck/templates"
)

func handleApp(w http.ResponseWriter, r *http.Request) {
	templates.Index("this is a test").Render(r.Context(), w)
}

func handleSwapLogin(w http.ResponseWriter, r *http.Request){
	
	templates.Login(templates.LOGIN_PARAMS, templates.LoginError{}).Render(r.Context(), w)
}

func handleSwapCreateAccount(w http.ResponseWriter, r *http.Request){
	
	templates.Login(templates.CREATE_USER_PARAMS, templates.LoginError{}).Render(r.Context(), w)
}

