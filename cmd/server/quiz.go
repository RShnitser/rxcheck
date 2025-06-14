package main

import(
	"net/http"
	"rxcheck/internal/auth"
	//"rxcheck/internal/database"
	"rxcheck/templates"
	//"fmt"
)

func (cfg *config)handleCreateQuiz(w http.ResponseWriter, r *http.Request){

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		return
	}

	// userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	// if err != nil {
	// 	//respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
	// 	return
	// }

	classification, err := cfg.db.GetClassificationByName(r.Context(), r.PathValue("drugClassification"))
	if err != nil{
		return
	}
	questions, err := cfg.db.ListRandomQuestionsByClassification(r.Context(), classification.ID)
	if err != nil{
		return
	}

	templates.Quiz(questions).Render(r.Context(), w)
}