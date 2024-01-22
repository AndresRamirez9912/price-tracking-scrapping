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
	router.Get("/scraping", handlers.ScrappingHandler)
	router.Get("/health", handlers.HealthHandler)

	// Start Server
	log.Println("Starting server at :3000 port")
	http.ListenAndServe(":3000", router)

}
