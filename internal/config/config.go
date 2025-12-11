package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Environment string
	Host        string
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        GetEnv("PORT", "8080"),
		DatabaseURL: GetEnv("DATABASE_URL", "./app.db"),
		Environment: GetEnv("ENVIRONMENT", "development"),
		Host:        GetEnv("HOST", "localhost"),
	}
}
