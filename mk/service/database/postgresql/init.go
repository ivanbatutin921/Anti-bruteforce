package postgresql

import (
	"log"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/config"
)

var DBDB = &PostgreSQLDB{}

func Init(cfg config.BruteForceConfig) (*PostgreSQLDB, error) {
	var err error
	err = DBDB.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := DBDB.Migrations(); err != nil {
		return nil, err
	}

	return DBDB, nil
}
