package main

import (
	"log"
	"net/http"

	"redis-auth/config"
	"redis-auth/handlers"
)

func main() {
	// Load configuration using Viper.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize Redis client.
	if err := config.InitRedis(cfg); err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}

	// Set up HTTP routes.
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/login", handlers.Login)

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
