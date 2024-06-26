package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sanjayheaven/ggb/internal/pkg/firebase"
	"github.com/sanjayheaven/ggb/internal/router"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Firebase
	firebase.InitFirebase()

	// Setup router
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}