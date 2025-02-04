package main

import(
	"net/http"
	"fmt"
)

func(cfg *config)handleAddUser(w http.ResponseWriter, r *http.Request){
	

	fmt.Println("Adding user")
	userName := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("name: %s, pass: %s\n", userName, password)
}