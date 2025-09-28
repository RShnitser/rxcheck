package main

import(
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"fmt"
	"os"
	"context"
	"encoding/json"
	"rxcheck/internal/database"
	"rxcheck/internal/auth"
	"database/sql"

	"github.com/joho/godotenv"
	"github.com/google/uuid"
)

type userData struct{
	name string
	pass string
}

type drugData struct{
	genericName string
	brandName string
	classification string
}

type questionData struct {
	GenericName   string   `json:"generic_name"`
	Question      string   `json:"question"`
	Choices       []string `json:"choices"`
	CorrectAnswer int64    `json:"correct_answer"`
	Explanation   string   `json:"explanation"`
}

const(
	ClassificationAnalgestic = "Analgesic"
	ClassificationNSAID = "NSAID"
	ClassificationStatin = "Statin"
)

func main(){
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}

	dbConnection, err := sql.Open("libsql", dbURL)
	if err != nil {
		fmt.Printf("Could not connect to database: %s\n", err)
		return
	}

	db := database.New(dbConnection)

	fmt.Println("seeding database")
	fmt.Println("seeding users")

	fmt.Println("deleting old users")
	err = db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("Could not delete users: %s\n", err)
		return
	}

	
	fmt.Println("creating test users")
	users := []userData{
		{"guest", "guest"},
	}

	for _, user := range users{
		hashPass, err := auth.HashPassword(user.pass)
		if err != nil{
			fmt.Println("unable to hash password")
			return
		}
		_, err = db.CreateUser(context.Background(), database.CreateUserParams{ uuid.New().String(), user.name, hashPass})
		if err != nil{
			fmt.Printf("unable to create user: %s\n", err)
			return
		}
	}

	fmt.Println("deleting drugs")
	err = db.DeleteDrugs(context.Background())
	if err != nil {
		fmt.Printf("Could not delete drugs: %s\n", err)
		return
	}


	fmt.Println("deleting classifications")
	err = db.DeleteClassifications(context.Background())
	if err != nil {
		fmt.Printf("Could not delete classifications: %s\n", err)
		return
	}


	fmt.Println("creating classifications")
	classifications := []string{
		ClassificationAnalgestic,
		ClassificationNSAID,
		ClassificationStatin,
	}

	for _, classification := range classifications{
		_, err = db.CreateClassification(context.Background(), database.CreateClassificationParams{uuid.New().String(), classification})
		if err != nil {
			fmt.Printf("Could not create classification: %s\n", err)
			return
		}
	}


	drugs := []drugData{
		{"acetaminophen", "tylenol", ClassificationAnalgestic},
		{"ibuprofen", "advil", ClassificationNSAID},
		{"atorvastatin", "lipitor", ClassificationStatin},
		{"rosuvastatin", "crestor", ClassificationStatin},
	}

	fmt.Println("creating drugs")

	for _, drug := range drugs{

		classification, err := db.GetClassificationByName(context.Background(), drug.classification)
		if err != nil {
			fmt.Printf("Could not get classification: %s\n", err)
			return
		}

		_, err = db.CreateDrug(context.Background(), database.CreateDrugParams{uuid.New().String(), drug.genericName, drug.brandName, classification.ID})
		if err != nil {
			fmt.Printf("Could not create drug: %s\n", err)
			return
		}
	}

	fmt.Println("deleting questions")
	err = db.DeleteQuestions(context.Background())
	if err != nil {
		fmt.Printf("Could not delete questions: %s\n", err)
		return
	} 

	questionsFile, err := os.Open("cmd/seed/questions.json")
	if err != nil{
		fmt.Printf("Could not open questions.json: %s\n", err)
		return
	}
	defer questionsFile.Close()

	var questions []questionData
	decoder := json.NewDecoder(questionsFile)
	err = decoder.Decode(&questions)
	if err != nil{
		fmt.Printf("Could not decode questions.json: %s\n", err)
		return
	}

	fmt.Println("creating questions")
	for _, qData := range questions{

		if len(qData.Choices) != 4{
			fmt.Printf("Question must have 4 answer choices: %s\n", qData.Question)
			continue
		}

		drug, err := db.GetDrugByGenericName(context.Background(), qData.GenericName)
		if err != nil{
			fmt.Printf("Could not find drug: %s\n", err)
			continue
		}

		params := database.CreateQuestionParams{
			uuid.New().String(),
			drug.ClassificationID,
			drug.ID,
			qData.Question,
			qData.Choices[0],
			qData.Choices[1],
			qData.Choices[2],
			qData.Choices[3],
			qData.Explanation,
			qData.CorrectAnswer,
		}

		_, err = db.CreateQuestion(context.Background(), params)
		if err != nil{
			fmt.Printf("Could not create questions: %s\n", err)
			continue
		}
	}
}