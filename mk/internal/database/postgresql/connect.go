package postgresql

import (
	"errors"
	"fmt"
	"log"

	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/config"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLDB struct {
	db *gorm.DB
}

func (db *PostgreSQLDB) Connect(cfg config.BruteForceConfig) error {
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

	log.Println("Успешное подключение к базе данных PostgreSQL")
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

func CheckLogin(db *PostgreSQLDB, user *models.Auth) error {
	var auth models.Auth
	db.db.Where("login = ?", user.Login).First(&auth)
	if auth.ID != 0 {
		return errors.New("пользователь уже существует")
	}
	return nil
}

func CheckIp(db *PostgreSQLDB, ip string) (bool, error) {
    err := db.db.Where("ip = ?", ip).First(&models.BlackList{}).Error
    if err == gorm.ErrRecordNotFound {
        return false, err
    }
    return true, nil
}

func CreateUser(db *PostgreSQLDB, user *models.Auth) error {
	err := db.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlackList(db *PostgreSQLDB, ip string) error {
	err := db.db.Where("ip = ?", ip).Delete(&models.BlackList{}).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateBlackList(db *PostgreSQLDB, bl *models.BlackList) error {
	blackList := models.BlackList{
		Ip: bl.Ip,
	}
	err := db.db.Create(&blackList).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateWhiteList(db *PostgreSQLDB, wl *models.WhiteList) error {
	whiteList := models.WhiteList{
		Ip: wl.Ip,
	}
	err := db.db.Create(&whiteList).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteWhiteList(db *PostgreSQLDB, ip string) error {
	err := db.db.Where("ip = ?", ip).Delete(&models.WhiteList{}).Error
	if err != nil {
		return err
	}
	return nil
}
