package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func SetRoutes() {
	http.HandleFunc("/api/verve/accept", AcceptRequest)
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	// Get query params
	id := r.URL.Query().Get("id")
	endpoint := r.URL.Query().Get("endpoint")

	// Send HTTP POST request if an endpoint is provided
	if endpoint != "" {
		SendPostRequest(endpoint, id)
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SendPostRequest(endpoint, id string) {
	url := fmt.Sprintf("%s?unique_id=%s", endpoint, id)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Failed to send POST request: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("POST request to %s with status code %d", endpoint, resp.StatusCode)
}
