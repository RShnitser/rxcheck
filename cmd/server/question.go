package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
	"strconv"
	//"github.com/google/uuid"
	//"fmt"
)

func getAnswerTextFromIndex(question database.Question, answer int64)string{
	switch answer{
	case 0:
		return question.Choice1
	case 1:
		return question.Choice2
	case 2:
		return question.Choice3
	case 3:
		return question.Choice4
	}

	return ""
}

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

	var questionID string
	var nextQuestionID string
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

	if int64(answer) == -1{
		//fmt.Println("Previous page was explanation.  Displaying question.")
		if session.QuestionIndex == 4{
			//fmt.Println("displaying summary")
			templates.Summary(session.Score).Render(r.Context(), w)
			return
		}
		templates.Question(question, session.QuestionIndex + 1).Render(r.Context(), w)
		return
	}
	
	newScore := session.Score
	if int64(answer) == question.AnswerIndex{
		newScore += 1
	}
		
	err = cfg.db.UpdateSession(r.Context(), database.UpdateSessionParams{newScore, session.QuestionIndex + 1, session.ID})
	if err != nil{
		//fmt.Println("could not update session")
		return
	}
	
	if int64(answer) != question.AnswerIndex{
		//fmt.Println("displaying explanation")
		correct := getAnswerTextFromIndex(question, question.AnswerIndex)
		incorrect := getAnswerTextFromIndex(question, int64(answer))
		templates.Explanation(question.Text, question.Explanation, correct, incorrect).Render(r.Context(), w)
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
	
	templates.Question(nextQuestion, session.QuestionIndex + 2).Render(r.Context(), w)
}