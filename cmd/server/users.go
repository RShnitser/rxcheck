package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	//"time"
	//"fmt"
	//"github.com/google/uuid"
	//"encoding/json"
	"rxcheck/templates"
)

func(cfg *config) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	password := r.FormValue("password")

	errs := templates.LoginError{
	}

	if userName == ""{
		errs.Name = "Enter a Username"
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	if password == ""{
		errs.Password = "Enter a password"
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	hashPass, err := auth.HashPassword(password)
	if err != nil{
		errs.General = "Server Error"
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	userParams := database.CreateUserParams{
		UserName: userName,
		HashedPassword: hashPass,
	}

	_, err = cfg.db.CreateUser(r.Context(), userParams)
	if err != nil {
		errs.General = "User Exists"
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		//respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	classifications, err := cfg.db.ListClassifications(r.Context())
	if err != nil{
		errs.General = "Server Error"
		templates.Login(templates.CREATE_USER_PARAMS, errs).Render(r.Context(), w)
		return
	}

	templates.Game(classifications).Render(r.Context(), w)
}
