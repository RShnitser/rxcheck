package main

import(
	"net/http"
	//"rxcheck/internal/database"
	"rxcheck/templates"
)

func (cfg *config)handleCreateQuiz(w http.ResponseWriter, r *http.Request){
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