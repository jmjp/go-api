package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
)

func startupVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to loading .env file")
	}
}
