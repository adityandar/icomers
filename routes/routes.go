package routes

import (
	"icomers/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// public routes
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// authenticated routes
	authRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes.Use(handlers.AuthMiddleware)

	// products route (subroute, authenticated)
	authRoutes.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	authRoutes.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
	authRoutes.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	authRoutes.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	authRoutes.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	return router
}
