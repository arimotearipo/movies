package main

import (
	"encoding/json"
	"fmt"

	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/arimotearipo/movies/internal/server"
)

func main() {

	dbConfig := initConfig()

	configJson, _ := json.MarshalIndent(dbConfig, "", "\t")

	fmt.Println("dbConfig", string(configJson))

	// Initialize database and perform connection
	db := database.NewDatabase(dbConfig)
	db.ConnectDB()

	// Initialize schema for database
	db.CreateSchemas()

	defer db.CloseDB()

	// Create new instance of Postgres Service layer
	store := psqlstorage.NewStorage(db.DB)

	port := getEnv("PORT", "3000")

	// Create server using the created Postgres Service layer
	server := server.NewServer(":"+port, store)
	server.Serve()
}
