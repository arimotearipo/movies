package main

import (
	"encoding/json"
	"fmt"

	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/arimotearipo/movies/internal/server"
	_ "github.com/lib/pq"
)

func main() {

	dbConfig := initConfig()

	configJson, _ := json.MarshalIndent(dbConfig, "", "\t")

	fmt.Println("dbConfig", string(configJson))

	db := database.NewDatabase(dbConfig)
	db.ConnectDB()

	db.CreateSchemas()

	defer db.CloseDB()

	store := psqlstorage.NewStorage(db.DB)

	port := getEnv("PORT", "8080")
	server := server.NewServer(":"+port, store)
	server.Serve()
}
