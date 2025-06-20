package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
	"strconv"
)

func (cfg *config)handleGetQuestion(w http.ResponseWriter, r *http.Request){
	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		return
	}

	nextQuestionIndex, err := strconv.Atoi(r.PathValue("nextQuestionIndex"))
	if err != nil{
		return
	}
}