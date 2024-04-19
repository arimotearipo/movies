package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDatabase(myEnv map[string]string) {
	// Form connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		myEnv["DB_HOST"], myEnv["DB_PORT"], myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_NAME"])

	var err error
	DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Connected to database")
}

func CloseDB() error {
	if DB == nil {
		log.Print("Database doesn't exist")
		return nil
	}

	err := DB.Close()
	if err != nil {
		log.Print("Error closing database connection", err)
		return err
	}
	log.Println("Connection to database closed")
	return nil
}
