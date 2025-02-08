package handlers

import "net/http"

func SetRoutes() {
	http.HandleFunc("/api/verve/accept", AcceptRequest)
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {

}

func SendPostRequest(w http.ResponseWriter, r *http.Request) {

}
