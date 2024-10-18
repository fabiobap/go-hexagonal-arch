package main

import (
	"log"

	"github.com/go-hexagonal-arch/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting the application")
	app.Start()
}
