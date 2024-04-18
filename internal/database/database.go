package database

import (
	"fmt"

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
		panic(err)
	}

	fmt.Println("Connected to database!")
}
