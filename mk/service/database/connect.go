package database

import (
	"time"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/pkg/logger"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DBDB *gorm.DB

func ConnectDb(url string, log *logger.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  url,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	log.Info("✅ Database connection successfully")

	log.Info("📦 Setting database connection pool...")
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
