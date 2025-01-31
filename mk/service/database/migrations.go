package database

import (
	
	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/models"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&models.BlackList{},
			&models.WhiteList{},
			&models.Auth{},
		}

		log.Info("📦 Creating types...")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}