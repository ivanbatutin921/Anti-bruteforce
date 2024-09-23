package postgresql

import (
	"github.com/ivanbatutin921/Anti-bruteforce/internal/config"
)

func Init(cfg config.BruteForceConfig) (*PostgreSQLDB, error) {
	db := &PostgreSQLDB{}
	err := db.Connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := db.Migrations(); err != nil {
		return nil, err
	}

	return db, nil
}
