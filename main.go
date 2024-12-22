package main

import (
	"fmt"
	"icomers/database"
	"icomers/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database.Init()
	// Initialize router
	router := routes.InitializeRoutes()

	// Start the server
	fmt.Println("Server running on http://localhost:8037")
	log.Fatal(http.ListenAndServe(":8037", router))

}
