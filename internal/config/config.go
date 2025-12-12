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
		Port:        GetEnv("PORT", "3000"),
		DatabaseURL: GetEnv("DATABASE_URL", "unknown"),
		Environment: GetEnv("ENVIRONMENT", "development"),
	}
}
