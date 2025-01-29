package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type BruteForceConfig struct{
	DBConfig
}

type DBConfig struct {
	DatabaseUrl  string
}

type HTTPCOnfig struct{}

type GRPCCongif struct{}

var Cfg BruteForceConfig

func LoadEnvVars() BruteForceConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(err.Error())
	}

	Cfg.DatabaseUrl = os.Getenv("DATABASE_URL")

	return Cfg

}
