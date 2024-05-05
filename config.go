package main

import (
	"log"
	"os"

	"github.com/arimotearipo/movies/internal/types"
	"github.com/joho/godotenv"
)

func initConfig() types.DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Couldn't load from .env file", err)
	}

	return types.DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "1234"),
		DBName:   getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return envValue
	}

	return fallback
}
