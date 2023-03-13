package main

import (
	"LearnECHO/internal/config"
	"LearnECHO/internal/database/postgres"
	"LearnECHO/internal/handlers"
	"LearnECHO/internal/router"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

func main() {
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed when parsing config: %v", err)
	}

	connectDB, err := postgres.ConnectDB(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = postgres.Migrate(connectDB)
	if err != nil {
		log.Fatalf("failed to run database migration: %v", err)
	}

	todoListHandler := handlers.NewTodoListHandlerImpl(connectDB)

	e := router.NewRouter(todoListHandler)

	e.Logger.Fatal(e.Start(":1234"))
	//e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort)))
}
