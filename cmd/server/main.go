package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Handle GET request
		fmt.Fprintf(w, "Handling GET request\n")
	case http.MethodPost:
		// Handle POST request
		fmt.Fprintf(w, "Handling POST request\n")
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
