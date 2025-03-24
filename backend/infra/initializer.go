package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not loaded: %v", err)
	}
}
