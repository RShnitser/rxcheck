package main

import(
	"net/http"
	"rxcheck/templates"
)

func handleApp(w http.ResponseWriter, r *http.Request) {
		templates.Index("this is a test").Render(r.Context(), w)
	}