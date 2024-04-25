package database

import (
	"fmt"
	"log"

	"github.com/arimotearipo/movies/internal/types"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB  *sqlx.DB
	cfg types.DBConfig
}

func NewDatabase(cfg types.DBConfig) *Database {
	return &Database{
		nil, cfg,
	}
}

func (db *Database) ConnectDB() {
	// Form connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		db.cfg.Host,
		db.cfg.Port,
		db.cfg.User,
		db.cfg.Password,
		db.cfg.DBName,
	)

	var err error
	db.DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Connected to database")
}

func (db *Database) CloseDB() error {
	if db.DB == nil {
		log.Print("Database doesn't exist")
		return nil
	}

	err := db.DB.Close()
	if err != nil {
		log.Print("Error closing database connection", err)
		return err
	}
	log.Println("Connection to database closed")
	return nil
}
