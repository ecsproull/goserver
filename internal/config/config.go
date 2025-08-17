package config

import (
	"os"
)

type Config struct {
	Port              string
	DatabaseURL       string
	JWTSecret         string
	MongoDatabase     string
	SendGridAPIKey    string
	SendGridFromEmail string
	FrontendURL       string
	GO_ENV            string
}

func Load() *Config {
	return &Config{
		Port:              getEnv("PORT", "3003"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key"),
		MongoDatabase:     getEnv("MONGO_DATABASE", "edandlinda"),
		SendGridAPIKey:    getEnv("SENDGRID_API_KEY", ""),
		SendGridFromEmail: getEnv("SENDGRID_FROM_EMAIL", ""),
		FrontendURL:       getEnv("FRONTEND_URL", "http://localhost:3001"),
		GO_ENV:            getEnv("GO_ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
