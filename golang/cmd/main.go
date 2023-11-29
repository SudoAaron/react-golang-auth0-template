package main

import (
	"log"
	"net/http"

	"github.com/SudoAaron/react-golang-auth0-template/api/private"
	"github.com/SudoAaron/react-golang-auth0-template/api/public"
	customMiddleware "github.com/SudoAaron/react-golang-auth0-template/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/public", public.Handler)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.EnsureValidToken)
		r.Get("/private", private.Handler)
	})

	log.Print("Server listening on http://localhost:3001/")
	if err := http.ListenAndServe("0.0.0.0:3001", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
