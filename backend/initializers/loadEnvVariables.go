package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load() // loads values from .env into the system

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
