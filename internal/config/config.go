package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	LoadEnvVars()
}

func LoadEnvVars() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf(err.Error())
	}
}
