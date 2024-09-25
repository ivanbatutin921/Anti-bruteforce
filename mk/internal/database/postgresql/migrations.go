package postgresql

import (
	"log"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
)

func (db *PostgreSQLDB) Migrations() error {
	// Implement your migrations logic here
	err := db.db.AutoMigrate(&models.Auth{}, &models.WhiteList{}, &models.BlackList{})
	if err != nil {
		return err
	}
	log.Println("Миграции выполнены")

	return nil
}
