package main

import (
	"AuthService/config"
)

func main() {
	server, err := InitializeServer()
	log := server.Logger
	if err != nil {
		log.Error("Failed to initialize server: " + err.Error())
		return
	}

	// Start the server
	envConfig := config.GetConfig()
	port := envConfig.AppPort
	if port == "" {
		port = "8080"
		log.Info("Defaulting to port " + port)
	}

	log.Info("Listening on port " + port)
	router := server.Router
	if err := router.Serve(":" + port); err != nil {
		log.Info("could not start server: " + err.Error())
	}
}
