package main

import (
	"RESTfulAPI-TodoList/internal/database/postgres"
	"RESTfulAPI-TodoList/internal/handlers"
	"RESTfulAPI-TodoList/internal/router"
	"RESTfulAPI-TodoList/models/config"
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

	var srv config.Server
	err = envconfig.Process("", &srv)
	if err != nil {
		log.Fatalf("failed when parsing server config: %v", err)
	}

	e.Logger.Fatal(e.Start(srv.Address))
}
