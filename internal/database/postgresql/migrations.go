package postgresql

import (
	"fmt"
	"reflect"
	model "root/internal/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	//db.Set("gorm:prepareStmt", false)

	models := []interface{}{
		&model.Auth{},
	}

	var errs []error
	for _, model := range models {
		if !db.Migrator().HasTable(model) {
			tableName := db.NamingStrategy.TableName(reflect.TypeOf(model).Elem().Name())
			stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (...)", tableName)
			err := db.Exec(stmt).Error
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return nil
}
