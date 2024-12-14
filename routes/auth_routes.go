package routes

import (
	"icomers/handlers"

	"github.com/gorilla/mux"
)

// public routes
func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
}
