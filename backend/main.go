package main

import (
	"log"
	"net/http"
	"product-management-backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define Product Routes
	r.HandleFunc("/products", handlers.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", handlers.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)

	// Start the server
	const port = ":8080"
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}




