package cmd

import (
	"fmt"
	"log"
	"net/http"

	"LoadBalancedAPI/cmd/config"
	"LoadBalancedAPI/service/handlers"
)

func main() {
	configuration, err := config.LoadConfig("env/properties.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	handlers.SetRoutes()

	// Start the server on port 8080
	log.Printf("Starting server on :%s", configuration.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", configuration.Host, configuration.Port), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
