package main

import (
	"fmt"
	"icomers/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize router
	router := routes.InitializeRoutes()

	// Start the server
	fmt.Println("Server running on http://localhost:8037")
	log.Fatal(http.ListenAndServe(":8037", router))

}
