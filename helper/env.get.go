package helper

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GetEnvValue(key string) string {
	err := godotenv.Load(".env")
	if (err != nil) {
		log.Fatal("Error load .env file")
	}

	return os.Getenv(key)
}