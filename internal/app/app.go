package app

import(
	config "github.com/ivanbatutin921/Anti-bruteforce/internal/config"
	db "github.com/ivanbatutin921/Anti-bruteforce/internal/database/postgresql"

	
)

func Run() {
	config.LoadEnvVars()
	db.Init(config.Cfg)	
}