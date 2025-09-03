package config

import (
	"os"
)

type Config struct {
	Port              string
	JWTSecret         string
	SendGridAPIKey    string
	SendGridFromEmail string
	FrontendURL       string
	GO_ENV            string
}

func Load() *Config {
	return &Config{
		Port:              getEnv("PORT", "3003"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key"),
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
