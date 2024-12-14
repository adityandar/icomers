package routes

import (
	"icomers/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	secure := r.PathPrefix("/api").Subrouter()
	secure.Use(handlers.AuthMiddleware)

	// products route (subroute, authenticated)
	secure.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	secure.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	secure.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	secure.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	secure.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
}
