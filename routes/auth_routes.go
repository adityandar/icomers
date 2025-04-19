package routes

import (
	"icomers/handlers"

	"github.com/gorilla/mux"
)

// public routes
func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/registerUser", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/loginUser", handlers.LoginUser).Methods("POST")
}
