package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/Carlosam7/calculator/internal/interfaces/http/handlers"
	"github.com/Carlosam7/calculator/internal/interfaces/http/routes"
)

func main() {
	healtHandler := handlers.NewHealthHandler()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	router := routes.SetupRoutes(healtHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: c.Handler(router),
	}

	log.Printf("Server starting on port %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
