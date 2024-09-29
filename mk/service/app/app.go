package app

import(
	config "github.com/ivanbatutin921/Anti-bruteforce/mk/service/config"
	db "github.com/ivanbatutin921/Anti-bruteforce/mk/service/database/postgresql"
	grpc "github.com/ivanbatutin921/Anti-bruteforce/mk/service/services/grpc"
	
	
)

func App() {
	config.LoadEnvVars()
	db.Init(config.Cfg)	
	grpc.ListenGRPC()
}