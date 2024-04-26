package main

import (
	"log"

	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/arimotearipo/movies/internal/server"
	"github.com/arimotearipo/movies/internal/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := types.DBConfig{
		Host:     myEnv["DB_HOST"],
		Port:     myEnv["DB_PORT"],
		User:     myEnv["DB_USER"],
		Password: myEnv["DB_PASSWORD"],
		DBName:   myEnv["DB_NAME"],
	}

	db := database.NewDatabase(dbConfig)
	db.ConnectDB()

	defer db.CloseDB()

	store := psqlstorage.NewStorage(db.DB)

	server := server.NewServer("localhost:"+myEnv["PORT"], store)
	server.Serve()
}
