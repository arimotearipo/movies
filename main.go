package main

import (
	"flag"

	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/arimotearipo/movies/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	createTable := flag.Bool(
		"createtable",
		false,
		"Set to true if your database is currently empty and you want this programme to create the tables for you",
	)
	flag.Parse()

	dbConfig := initConfig()

	db := database.NewDatabase(dbConfig)
	db.ConnectDB()

	if *createTable {
		db.CreateSchemas()
	}

	defer db.CloseDB()

	store := psqlstorage.NewStorage(db.DB)

	port := getEnv("PORT", "8080")
	server := server.NewServer("localhost:"+port, store)
	server.Serve()
}
