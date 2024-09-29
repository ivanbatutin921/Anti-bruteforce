package app

import(
	config "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/config"
	db "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/database/postgresql"
	grpc "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/services/grpc"
	
	
)

func Run() {
	config.LoadEnvVars()
	db.Init(config.Cfg)	
	grpc.ListenGRPC()
}