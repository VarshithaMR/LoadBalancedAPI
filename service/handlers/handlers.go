package handlers

import (
	"net/http"
)

func SetRoutes() {
	http.HandleFunc("/api/verve/accept", AcceptRequest)
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SendPostRequest(w http.ResponseWriter, r *http.Request) {

}
