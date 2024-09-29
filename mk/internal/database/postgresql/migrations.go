package postgresql

import (
	"log"
	"time"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
)

func (db *PostgreSQLDB) Migrations() error {
	// Implement your migrations logic here
	start:=time.Now()
	err := db.db.AutoMigrate(&models.Auth{}, &models.WhiteList{}, &models.BlackList{})
	if err != nil {
		return err
	}
	elapsed := time.Since(start)

	log.Println("Миграции выполнены Время: ", elapsed)
	
	return nil
}
