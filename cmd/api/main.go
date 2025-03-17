package main

import (
	"log"
	"net/http"

	"product-api/internal/config"
	"product-api/internal/db"
	"product-api/internal/handlers"
	"product-api/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := db.NewDatabase(cfg.DBConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Set up router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", handlers.Register(database)).Methods("POST")
	r.HandleFunc("/login", handlers.Login(database)).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/products").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("", handlers.GetProducts(database)).Methods("GET")
	protected.HandleFunc("", handlers.CreateProduct(database)).Methods("POST")
	protected.HandleFunc("/{id}", handlers.GetProduct(database)).Methods("GET")
	protected.HandleFunc("/{id}", handlers.UpdateProduct(database)).Methods("PUT")
	protected.HandleFunc("/{id}", handlers.DeleteProduct(database)).Methods("DELETE")

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
