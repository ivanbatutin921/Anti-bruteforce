package postgresql

import (
	"fmt"
	"log"
	"time"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/config"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLDB struct {
	db *gorm.DB
}

func (db *PostgreSQLDB) Connect(cfg config.BruteForceConfig) error {
	start := time.Now()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PGHOST, cfg.PGUSER, cfg.PGPASSWORD, cfg.PGDATABASE, cfg.PGPORT)

	var err error
	db.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Println("Не удалось подключиться к базе данных PostgreSQL:", err)
		return err
	}

	log.Println(db.db)

	elapsed := time.Since(start)
	log.Println("Успешное подключение к базе данных PostgreSQL. Время: ", elapsed)
	return nil
}

func (db *PostgreSQLDB) Close() {
	sqlDB, err := db.db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println(err)
	}
}

func (db *PostgreSQLDB) CheckLogin(user *models.Auth) (*models.Auth, error) {
	var auth models.Auth
	db.db.Where("login = ?", user.Login).First(&auth)
	if auth.ID != 0 {
		return &auth, nil
	}
	return nil, nil // no error if user does not exist
}

func (db *PostgreSQLDB) CheckIp(ip string) bool {
	if db.db == nil {
		log.Println("db object is not initialized")
		return false
	}
	err := db.db.Where("ip = ?", ip).First(&models.BlackList{}).Error
	if err == gorm.ErrRecordNotFound {
		log.Println(err)
		return false
	}
	return true
}

func (db *PostgreSQLDB) CreateUser(user *models.Auth) error {
	err := db.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *PostgreSQLDB) DeleteBlackList(ip string) error {
	err := db.db.Where("ip = ?", ip).Delete(&models.BlackList{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *PostgreSQLDB) CreateBlackList(bl *models.BlackList) error {
	blackList := models.BlackList{
		Ip: bl.Ip,
	}
	err := db.db.Create(&blackList).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *PostgreSQLDB) CreateWhiteList(wl *models.WhiteList) error {
	whiteList := models.WhiteList{
		Ip: wl.Ip,
	}
	err := db.db.Create(&whiteList).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *PostgreSQLDB) DeleteWhiteList(ip string) error {
	err := db.db.Where("ip = ?", ip).Delete(&models.WhiteList{}).Error
	if err != nil {
		return err
	}
	return nil
}
