package main

import (
	"fmt"
	"log"
	"net/http"

	"LoadBalancedAPI/cmd/config"
	"LoadBalancedAPI/service/handlers"
	"LoadBalancedAPI/service/persistence"
)

func main() {
	configuration, err := config.LoadConfig("env/properties.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	persistence.InitRedis(configuration.Host, configuration.RedisPort)
	handlers.SetRoutes()

	// Start the server on port 8080
	log.Printf("Starting server on :%d", configuration.Port)
	if err = http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Host, configuration.Port), nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
