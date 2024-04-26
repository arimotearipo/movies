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

func (db *Database) checkSchemaExists() bool {
	if db.DB == nil {
		log.Fatal("Database connection doesn't exist")
	}

	queryString := `---sql
	SELECT 
    EXISTS (
        SELECT 1 
        FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'directors'
    ) AS table1_exists,
    EXISTS (
        SELECT 1 
        FROM information_schema.tables 
        WHERE table_schema = 'public' 
        AND table_name = 'movies'
    ) AS table2_exists;
	`

	var result []types.SchemaExists
	err := db.DB.Select(&result, queryString)
	if err != nil {
		log.Fatal("Error querying tables")
	}

	if !result[0].T1 || !result[0].T2 {
		return false
	}

	return true
}

func (db *Database) CreateSchemas() {
	if db.DB == nil {
		log.Fatal("Database connection doesn't exists")
	}

	exists := db.checkSchemaExists()
	if exists {
		log.Fatal("Schemas already exists")
	}

	queryString := `---sql
		CREATE TABLE directors (
			director_id SERIAL PRIMARY KEY,
			name	VARCHAR(80),
			date_of_birth TIMESTAMP,
			gender CHAR(1),
			nationality VARCHAR(20)
		);
		CREATE TABLE movies (
			movie_id SERIAL PRIMARY KEY,
			title	VARCHAR(80),
			director_id INTEGER,
			year INTEGER,
			FOREIGN KEY (director_id) REFERENCES public.directors(director_id)
		);
	`

	_, err := db.DB.Exec(queryString)
	if err != nil {
		log.Fatal("Error creating tables")
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
