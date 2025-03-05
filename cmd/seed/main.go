package main

import(
	_ "github.com/lib/pq"
	"fmt"
	"os"
	"context"
	"rxcheck/internal/database"
	"rxcheck/internal/auth"
	"database/sql"

	"github.com/joho/godotenv"
)

type userData struct{
	name string
	pass string
}

func main(){
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL not found")
		return
	}

	dbConnection, err := sql.Open("postgres", dbURL)
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
		_, err = db.CreateUser(context.Background(), database.CreateUserParams{user.name, hashPass})
		if err != nil{
			fmt.Println("unable to create user")
			return
		}
	}

	err = db.DeleteClassifications(context.Background())
	if err != nil {
		fmt.Printf("Could not delete classifications: %s\n", err)
		return
	}

	classifications := []string{
		"Analgesic",
	}

	for _, classification := range classifications{
		_, err = db.CreateClassification(context.Background(), classification)
		if err != nil {
			fmt.Printf("Could not create classification: %s\n", err)
			return
		}
	}
}