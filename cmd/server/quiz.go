package main

import(
	"net/http"
	"rxcheck/internal/auth"
	"rxcheck/internal/database"
	"rxcheck/templates"
	//"fmt"

	"github.com/google/uuid"
)

func (cfg *config)handleCreateQuiz(w http.ResponseWriter, r *http.Request){

	token, err := auth.GetBearerToken(r.Header)
	if err != nil{
		//fmt.Println("No Bearer token")
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		//respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		//fmt.Println(err)
		return
	}

	err = cfg.db.DeleteSession(r.Context(), userID)
	if err != nil{
		//fmt.Println("could not delete")
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
	//fmt.Printf("Question Length: %d", len(questions))

	sessionParams := database.CreateSessionParams{
		uuid.New().String(),
		userID,
		questions[0].ID,
		questions[1].ID,
		questions[2].ID,
		questions[3].ID,
		questions[4].ID,
	}
	_, err = cfg.db.CreateSession(r.Context(), sessionParams)
	if err != nil{
		return
	}

	templates.Question(questions[0], 1).Render(r.Context(), w)
}