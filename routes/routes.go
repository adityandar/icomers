package routes

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	RegisterAuthRoutes(router)

	RegisterProductRoutes(router)

	return router
}
