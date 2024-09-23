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
	PGHOST     string
	PGUSER     string
	PGPASSWORD string
	PGDATABASE string
	PGPORT     string
}

type HTTPCOnfig struct{}

type GRPCCongif struct{}

var Cfg BruteForceConfig

func LoadEnvVars() BruteForceConfig {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf(err.Error())
	}

	Cfg.PGHOST = os.Getenv("PGHOST")
	Cfg.PGUSER = os.Getenv("PGUSER")
	Cfg.PGPASSWORD = os.Getenv("PGPASSWORD")
	Cfg.PGDATABASE = os.Getenv("PGDATABASE")
	Cfg.PGPORT = os.Getenv("PGPORT")

	return Cfg

}
