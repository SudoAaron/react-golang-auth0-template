package main

import (
	"log"
	"net/http"

	"github.com/SudoAaron/react-golang-auth0-template/api/admin"
	"github.com/SudoAaron/react-golang-auth0-template/api/get_roles"
	"github.com/SudoAaron/react-golang-auth0-template/api/init_user"
	"github.com/SudoAaron/react-golang-auth0-template/api/private"
	"github.com/SudoAaron/react-golang-auth0-template/api/public"
	"github.com/SudoAaron/react-golang-auth0-template/api/set_role"
	"github.com/SudoAaron/react-golang-auth0-template/db"
	customMiddleware "github.com/SudoAaron/react-golang-auth0-template/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	// Set up DB connection
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
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
		r.Get("/init_user", init_user.Handler(dbConn))

		// No roles necessary for route
		r.Group(func(r chi.Router) {
			r.Get("/get_roles", get_roles.Handler(dbConn))
		})

		// Users only
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.CheckRoles([]string{"user"}, dbConn))
			r.Get("/set_role", set_role.Handler(dbConn))
		})

		// Require Admins only
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.CheckRoles([]string{"admin"}, dbConn))
			r.Get("/admin", admin.Handler)
		})

		// Require User or Admin Role
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.CheckRoles([]string{"user", "admin"}, dbConn))
			r.Get("/private", private.Handler)
		})

	})

	log.Print("Server listening on http://localhost:3001/")
	if err := http.ListenAndServe("0.0.0.0:3001", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
