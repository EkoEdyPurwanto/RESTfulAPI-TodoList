package main

import (
	"LearnECHO/internal/config"
	"LearnECHO/internal/database"
	"LearnECHO/internal/handlers"
	"LearnECHO/internal/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	log := e.Logger

	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed when parsing config: %v", err)
	}

	connectDB, err := database.ConnectDB(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	todoListHandler := handlers.NewTodoListHandlerImpl(connectDB, log)

	e = router.NewRouter(todoListHandler)

	e.Logger.Fatal(e.Start(":1234"))
}
