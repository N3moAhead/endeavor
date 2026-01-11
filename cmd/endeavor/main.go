package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/N3moAhead/endeavor/internal/db"
	"github.com/N3moAhead/endeavor/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	// --- Load environment variables ---
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("You have to set the env variable DATABASE_URL")
	}

	// --- Init Systems ---
	db.Init()
	router := router.New()

	// --- Application Start ---

	s := http.Server{
		Addr:           ":9090",
		Handler:        router.GetHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server listening on: http://localhost:9090")
	log.Fatal(s.ListenAndServe())
}
