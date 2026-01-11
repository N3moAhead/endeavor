package main

import (
	"log"
	"os"

	"github.com/N3moAhead/endeavor/internal/db"
	"github.com/N3moAhead/endeavor/internal/db/seed"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("You have to set the env variable DATABASE_URL")
	}

	// Initialize database
	db.Init()

	// Initialize seed with database connection
	seed.Init(db.DB)

	// Seed data
	if err := seed.SeedAll(); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	log.Println("Database seeded successfully")
}
