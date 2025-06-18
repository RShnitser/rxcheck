package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
	//"fmt"
)

func (cfg *config)handleCreateQuiz(w http.ResponseWriter, r *http.Request){

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		//respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	err = cfg.db.DeleteQuiz(r.Context(), userID)
	if err != nil{
		return
	}

	classification, err := cfg.db.GetClassificationByName(r.Context(), r.PathValue("drugClassification"))
	if err != nil{
		return
	}
	questions, err := cfg.db.ListRandomQuestionsByClassification(r.Context(), classification.ID)
	if err != nil{
		return
	}

	quizParams := database.CreateQuizParams{
		userID,
		questions[0].ID,
		questions[1].ID,
		questions[2].ID,
		questions[3].ID,
		questions[4].ID,
	}
	quiz, err := cfg.db.CreateQuiz(r.Context(), quizParams)
	if err != nil{
		return
	}

	templates.Question(questions[0], quiz.NextQuestionIndex).Render(r.Context(), w)
}