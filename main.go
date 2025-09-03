package main

import (
	"log"
	"os"

	"goserver/internal/config"
	"goserver/internal/database"
	"goserver/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// open (or create) log file, append mode
	f, err := os.OpenFile("goserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	// keep file open for lifetime of process (don't defer f.Close() here if server runs)
	log.SetOutput(f)

	log.Printf("server starting...")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	cfg := config.Load()

	// Connect to PostgreSQL database
	if err := database.ConnectDatabase(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		log.Println("Database connection established")
	}
	defer database.CloseDatabase()

	r := router.SetupRouter()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(r.Run(cfg.Port))

}
