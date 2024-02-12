package main

import (
	"log"
	"net/http"
	"os"
	"scrapper/src/constants"
	"scrapper/src/handlers"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Read the .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading the .env variables", err)
		return
	}

	// Router
	router := chi.NewRouter()

	// Endpoints
	router.Post("/api/scraping", handlers.ScrappingHandler)
	router.Get("/health", handlers.HealthHandler)

	// Start Server
	log.Println("Starting server at port", os.Getenv(constants.PORT))
	http.ListenAndServe(os.Getenv(constants.PORT), router)

}
