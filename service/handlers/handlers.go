package handlers

import (
	"fmt"
	"log"
	"net/http"

	"LoadBalancedAPI/service"
)

func SetRoutes() {
	http.HandleFunc("/api/verve/accept", AcceptRequest)
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	// Get query params
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
	}

	uniqueRequestCount := service.TrackUniqueRequest(id)
	endpoint := r.URL.Query().Get("endpoint")

	// Send HTTP POST request if an endpoint is provided
	if endpoint != "" {
		err := SendPostRequest(endpoint, uniqueRequestCount)
		if err != nil {
			log.Printf("Error sending POST request: %v", err)
			http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
			return
		}
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SendPostRequest(endpoint string, count int) error {
	url := fmt.Sprintf("%s?count=%d", endpoint, count)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Printf("Failed to send POST request: %v", err)
		return err
	}
	defer resp.Body.Close()

	log.Printf("POST request to %s with status code %d", endpoint, resp.StatusCode)
	return nil
}
