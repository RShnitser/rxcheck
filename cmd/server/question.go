package main

import(
	"net/http"
	"rxcheck/internal/auth"
	// "rxcheck/internal/database"
	"rxcheck/templates"
	"strconv"
	"github.com/google/uuid"
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

	quiz, err := cfg.db.GetQuizByUserID(r.Context(), userID)
	if err != nil{
		return
	}

	var questionID uuid.UUID
	switch nextQuestionIndex{
	case 0:
		questionID = quiz.Question1
	case 1:
		questionID = quiz.Question2
	case 2:
		questionID = quiz.Question3
	case 3:
		questionID = quiz.Question4
	case 5:
		questionID = quiz.Question5
	default:
		return
	}

	question, err := cfg.db.GetQuestionByID(r.Context(), questionID)
	if err != nil{
		return
	}

	if nextQuestionIndex >= 4{
		return
	}
	templates.Question(question, int32(nextQuestionIndex + 1)).Render(r.Context(), w)
}