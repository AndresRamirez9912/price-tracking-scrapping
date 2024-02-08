package main

import (
	"log"
	"net/http"
	"scrapper/src/handlers"

	"github.com/go-chi/chi"
)

func main() {
	// Router
	router := chi.NewRouter()

	// Endpoints
	router.Post("/api/scraping", handlers.ScrappingHandler)
	router.Get("/health", handlers.HealthHandler)

	// Start Server
	log.Println("Starting server at :3002 port")
	http.ListenAndServe(":3002", router)

}
