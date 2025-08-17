package main

import (
	"log"

	"goserver/internal/config"
	"goserver/internal/database"
	"goserver/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	cfg := config.Load()
	database.InitMongo(cfg.DatabaseURL)

	r := router.SetupRouter(cfg)

	port := cfg.Port
	if port == "" {
		port = "3003"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
