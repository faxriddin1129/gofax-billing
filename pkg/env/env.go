package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env variables...")
	}
}

func GetEnv(key string) string {
	val := os.Getenv(key)
	return val
}
