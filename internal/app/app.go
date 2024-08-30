package app

import(
	db "root/internal/database/postgresql"
	config "root/internal/config"
	
)

func Run() {
	db.Connect()
	db.MigrateModels(db.DB.DB)
	config.LoadEnvVars()
}