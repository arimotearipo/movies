package main

import (
	"os"

	"github.com/arimotearipo/movies/internal/types"
)

func initConfig() types.DBConfig {
	return types.DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return envValue
	}

	return fallback
}
