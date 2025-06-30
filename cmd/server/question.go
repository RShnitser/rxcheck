package main

import(
	"net/http"
	"rxcheck/internal/auth"
	// "rxcheck/internal/database"
	"rxcheck/templates"
	"strconv"
	"github.com/google/uuid"
	//"fmt"
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

	questionIndex, err := strconv.Atoi(r.PathValue("questionIndex"))
	if err != nil{
		return
	}

	//fmt.Printf("next question index %d\n", nextQuestionIndex)

	quiz, err := cfg.db.GetQuizByUserID(r.Context(), userID)
	if err != nil{
		return
	}

	var questionID uuid.UUID
	var nextQuestionID uuid.UUID
	switch questionIndex{
	case 0:
		questionID = quiz.Question1
		nextQuestionID = quiz.Question2
	case 1:
		questionID = quiz.Question2
		nextQuestionID = quiz.Question3
	case 2:
		questionID = quiz.Question3
		nextQuestionID = quiz.Question4
	case 3:
		questionID = quiz.Question4
		nextQuestionID = quiz.Question5
	case 4:
		questionID = quiz.Question5
	default:
		return
	}

	question, err := cfg.db.GetQuestionByID(r.Context(), questionID)
	if err != nil{
		return
	}

	answer, err := strconv.Atoi(r.FormValue("answer"))
	if err != nil{
		return
	}
	
	if int32(answer) != -1 && int32(answer) != question.AnswerIndex{
		templates.Explanation(question.Explanation, int32(questionIndex)).Render(r.Context(), w)
		return
	}

	if questionIndex == 4{
		templates.Summary().Render(r.Context(), w)
		return
	}
	
	nextQuestion, err := cfg.db.GetQuestionByID(r.Context(), nextQuestionID)
	if err != nil{
		//fmt.Println("Invalid Next Question ID")
		return
	}
	
	templates.Question(nextQuestion, int32(questionIndex + 1)).Render(r.Context(), w)
}