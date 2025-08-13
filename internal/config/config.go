package config

import (
	"os"
)

type Config struct {
	Port          string
	DatabaseURL   string
	JWTSecret     string
	MongoDatabase string
}

func Load() *Config {
	return &Config{
		Port:          getEnv("PORT", "3003"),
		DatabaseURL:   getEnv("DATABASE_URL", "mongodb://localhost:27017"),
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		MongoDatabase: getEnv("MONGO_DATABASE", "edandlinda"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
