package postgresql

import (
	"log"
	model "root/internal/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	//db.Set("gorm:prepareStmt", false)

	models := []interface{}{
		&model.Auth{},
		&model.WhiteList{},
		&model.BlackList{},
	}

	for _, model := range models {
		if !db.Migrator().HasTable(model) {
			log.Printf("миграция для модели %T", model)
			db.AutoMigrate(model)
		} else {
			log.Printf("Таблица для модели %T уже существует, миграция пропущена", model)
		}
	}
}
