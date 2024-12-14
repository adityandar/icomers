package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	router := mux.NewRouter()

	// define a basic route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Icomers3!")
	}).Methods("GET")

	// Start the server
	fmt.Println("Server running on http://localhost:8037")
	err := http.ListenAndServe(":8037", router)
	if err != nil {
		fmt.Println("error,", err)
	}

}
