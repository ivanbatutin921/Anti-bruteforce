package postgresql

import "gorm.io/gorm"

type DataBase interface {
	Connect() (*gorm.DB, error)
	Disconnect()
	Migrations () error
}
