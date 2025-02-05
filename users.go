package main

import(
	"net/http"
	"fmt"

	"rxcheck/internal/auth"
	"rxcheck/internal/database"
)

func(cfg *config)handleAddUser(w http.ResponseWriter, r *http.Request){
	

	fmt.Println("Adding user")
	userName := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("name: %s, pass: %s\n", userName, password)

	hash, err := auth.HashPassword(password)
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	_, err = cfg.db.CreateUser(r.Context(), database.CreateUserParams{userName, hash})
	if err != nil {
		//respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}
}