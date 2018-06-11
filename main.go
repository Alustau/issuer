package main

import (
	"fmt"
	"log"

	"github.com/alustau/issuer/router"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Print("Starting server")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Print("[OK]")
	router.Routes()
}
