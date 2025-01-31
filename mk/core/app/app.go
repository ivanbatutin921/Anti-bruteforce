package app

import (
	"fmt"

	database "github.com/ivanbatutin921/Anti-bruteforce/mk/service/database"
	grpc "github.com/ivanbatutin921/Anti-bruteforce/mk/service/services/grpc"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/pkg/config"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/pkg/logger"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	logger *logger.Logger
	db     *gorm.DB
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {

	err := app.initDeps()

	if err != nil {
		return err
	}

	app.runGrpcServer()

	return nil
}

func (app *App) initDeps() error {

	inits := []func() error{
		app.initConfig,
		app.initLogger,
		app.initDb,

	}
	for _, init := range inits {
		err := init()
		if err != nil {
			return fmt.Errorf("%s", "✖ Failed to initialize dependencies: "+err.Error())
		}
	}
	return nil
}

func (app *App) initConfig() error {
	if app.config == nil {
		config, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("%s", "✖ Failed to load config: "+err.Error())
		}
		app.config = config
	}

	err := config.Load()
	if err != nil {
		return fmt.Errorf("%s", "✖ Failed to load config: "+err.Error())
	}

	return nil
}

func (app *App) initLogger() error {
	if app.logger == nil {
		app.logger = logger.GetLogger()
	}
	return nil
}

func (app *App) initDb() error {
	if app.db == nil {
		db, err := database.ConnectDb(app.config.DatabaseUrl, app.logger)
		if err != nil {
			return err
		}
		app.db = db

		// true - запустить миграцию
		// false - не запускать
		if err := database.Migrate(db, false, app.logger); err != nil {
			return fmt.Errorf("%s", "✖ Failed to migrate database: "+err.Error())
		}
	}

	return nil
}

func(app *App) runGrpcServer() error {
	grpc.ListenGRPC()
	grpc.NewServer(app.db,app.logger)
	return nil
}