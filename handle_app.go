package main

import(
	"net/http"
	"rxcheck/templates"
)

func handleApp(w http.ResponseWriter, r *http.Request) {
	templates.Index("this is a test").Render(r.Context(), w)
}

func handleSwapLogin(w http.ResponseWriter, r *http.Request){
	params := templates.LoginParams{
		Title: "Login",
		SwapMessage: "Don't Have an account? Create Account",
		SubmitURL: "/login_user",
		SwapURL: "/swap_create",
	}
	templates.Login(params, templates.LoginError{}).Render(r.Context(), w)
}

func handleSwapCreateAccount(w http.ResponseWriter, r *http.Request){
	params := templates.LoginParams{
		Title: "Create User",
		SwapMessage: "Already have an account? Sign In",
		SubmitURL: "/create_user",
		SwapURL: "/swap_login",
	}
	templates.Login(params, templates.LoginError{}).Render(r.Context(), w)
}
