package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
	"strconv"
	"github.com/google/uuid"
	//"fmt"
)

func (cfg *config)handleGetQuestion(w http.ResponseWriter, r *http.Request){
	
	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		//fmt.Println("No token")
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		//fmt.Println("Invalid token")
		return
	}

	session, err := cfg.db.GetSessionByUserID(r.Context(), userID)
	if err != nil{
		//fmt.Println("Could not get session")
		return
	}

	var questionID uuid.UUID
	var nextQuestionID uuid.UUID
	switch session.QuestionIndex{
	case 0:
		questionID = session.Question1
		nextQuestionID = session.Question2
	case 1:
		questionID = session.Question2
		nextQuestionID = session.Question3
	case 2:
		questionID = session.Question3
		nextQuestionID = session.Question4
	case 3:
		questionID = session.Question4
		nextQuestionID = session.Question5
	case 4:
		questionID = session.Question5
	default:
		//fmt.Printf("invalid question index:%d\n", session.QuestionIndex)
		return
	}

	question, err := cfg.db.GetQuestionByID(r.Context(), questionID)
	if err != nil{
		//fmt.Println("could not get question")
		return
	}

	answer, err := strconv.Atoi(r.FormValue("answer"))
	if err != nil{
		//fmt.Println("could not convert answer")
		return
	}

	if int32(answer) == -1{
		//fmt.Println("Previous page was explanation.  Displaying question.")
		if session.QuestionIndex == 4{
			//fmt.Println("displaying summary")
			templates.Summary(session.Score).Render(r.Context(), w)
			return
		}
		templates.Question(question).Render(r.Context(), w)
		return
	}
	
	newScore := session.Score
	if int32(answer) == question.AnswerIndex{
		newScore += 1
	}
		
	err = cfg.db.UpdateSession(r.Context(), database.UpdateSessionParams{session.ID, newScore, session.QuestionIndex + 1})
	if err != nil{
		//fmt.Println("could not update session")
		return
	}
	
	if int32(answer) != question.AnswerIndex{
		//fmt.Println("displaying explanation")
		templates.Explanation(question.Explanation).Render(r.Context(), w)
		return
	}

	if session.QuestionIndex == 4{
		//fmt.Println("displaying summary")
		templates.Summary(newScore).Render(r.Context(), w)
		return
	}
	
	nextQuestion, err := cfg.db.GetQuestionByID(r.Context(), nextQuestionID)
	if err != nil{
		//fmt.Println("Invalid Next Question ID")
		return
	}
	
	templates.Question(nextQuestion).Render(r.Context(), w)
}