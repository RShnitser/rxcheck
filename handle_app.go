package main

import(
	"net/http"
	"rxcheck/templates"
)

func handleApp(w http.ResponseWriter, r *http.Request) {
		templates.Index("test").Render(r.Context(), w)
	}