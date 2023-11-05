package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file -->", err)
		return
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connection := fmt.Sprintf("host=localhost port=5432 user=%s password='%s' dbname=%s sslmode=disable", user, password, dbname)

	db, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Error connecting to database -->", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database -->", err)
		return
	}

	fmt.Println("Successfully connected to the database.")
}
