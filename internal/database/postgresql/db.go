package postgresql

import (
	"github.com/ivanbatutin921/Anti-bruteforce/internal/config"
	model "github.com/ivanbatutin921/Anti-bruteforce/internal/models"
	"gorm.io/gorm"
)

type DB interface {
	Connect(config.BruteForceConfig) (*gorm.DB, error)
	Close()
	Migrations() error
	CheckLogin(user *model.Auth) error
	CreateUser(user *model.Auth) error
}
