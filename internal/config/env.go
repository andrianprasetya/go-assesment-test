package config

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning Error loading .env file (%v)", err)
	}
}
